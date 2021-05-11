package controller

import (
	"errors"
	"fmt"
	"fpdapp/models/entity"
	"fpdapp/serializers/detail"
	"fpdapp/serializers/index"
	"fpdapp/validators/post_edit"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//*******************************************************************
// パラメータ定義
//*******************************************************************
//公開設定
const (
	private = "0"
	public  = "1"
)

//*******************************************************************
// 公開/非公開いずれかの投稿数を取得する
//*******************************************************************
func (cont *Controller) CountPost(c *gin.Context) {
	publishing := c.DefaultQuery("publishing", public)
	count := cont.DbConn.CountPost(publishing)
	response := struct {
		Count int `json:"count"`
	}{
		count,
	}
	c.JSON(http.StatusOK, response)
}

//*******************************************************************
// 投稿を1ページ表示件数分取得する
//*******************************************************************
func (cont *Controller) IndexList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	publishing := c.DefaultQuery("publishing", public)

	posts := cont.DbConn.FindIndex(page, publishing)
	var response []index.Response
	// レスポンスの形でリスト化
	for _, post := range posts {
		resp := index.Serializer{Post: post}
		response = append(response, resp.Response())
	}
	c.JSON(http.StatusOK, response)
}

//*******************************************************************
// 投稿を対象idの1件取得する
//*******************************************************************
func (cont *Controller) FetchPost(c *gin.Context) {
	postId := c.Query("postId")

	postModel := cont.DbConn.FindPost(postId)
	detailSerializer := detail.Serializer{Post: postModel}

	c.JSON(http.StatusOK, detailSerializer.Response())
}

//*******************************************************************
// 投稿記事テーブルへ記事を1件登録する
//*******************************************************************
func (cont *Controller) Create(c *gin.Context) {
	// JSONを構造体に置き換える（不要なjsonデータを受け取った場合、カットする）
	var request post_edit.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := post_edit.Validator{Post: request}
	// 構造体をPost型にバインドしてinsertする
	createdPostId := cont.DbConn.InsertPost(req.Request())
	c.JSON(http.StatusCreated, createdPostId)
}

//*******************************************************************
// クライアントから送信されたファイルをローカルに仮保存し、awsS3にアップロードし、
// S3の{オブジェクトURL}と{オブジェクトキー}をレスポンスとして返す。
//*******************************************************************
func (cont *Controller) ImageUpload(c *gin.Context) {
	userId := c.PostForm("user_id")              //投稿者のuserID
	file, err := c.FormFile("image")             //アップロード対象ファイル
	position := strToInt(c.PostForm("position")) //画像の順番
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	tmpDirName := "/app/images/"                   //S3にアップロードするファイルを仮置きするディレクトリ
	saveFileFullPath := tmpDirName + file.Filename // ex: /app/images/wan-chan.png

	// 仮置き用ディレクトリがなければ作成する
	if _, err := os.Stat(tmpDirName); os.IsNotExist(err) {
		os.Mkdir(tmpDirName, 0777)
	}

	// 受信したファイルを仮置き用ディレクトリへ保存する
	if err := c.SaveUploadedFile(file, saveFileFullPath); err != nil {
		fmt.Println(err.Error())
	}

	// 仮置きしたファイルをS3へアップロードする
	objectUrl, objectKey, err := UploadToS3(saveFileFullPath, userId)
	if err != nil {
		fmt.Println(err.Error())
	}

	// postImagesテーブルへオブジェクトキーを登録する（idは空）
	targetStruct := entity.PostImage{
		// FIXME: positionが0にならず、1になる
		Position: position,
		// TODO: post-postimage-imagesテーブル定義を確定後、処理を追加する
		//ImagePath: objectKey,
	}
	registeredPostImage := cont.DbConn.InsertPostImage(&targetStruct)

	// S3のオブジェクトURLとオブジェクトキーをjsonでレスポンスする
	response := struct {
		ObjectUrl string `json:"object_url"`
		ObjectKey string `json:"object_key"`
		Position  int    `json:"position"`
	}{
		objectUrl,
		objectKey,
		registeredPostImage.Position,
	}

	// {ObjectUrl:https://bbsapp-img.s3.us-east-2.amazonaws.com/images/1_b3dbd289-25c7-4f15-ad9d-46fd1f16f2ff.png ObjectKey:images/1_b3dbd289-25c7-4f15-ad9d-46fd1f16f2ff.png Position:1}
	c.JSON(http.StatusCreated, response)

	// 仮置きディレクトリに入っているファイルを削除する
	// TODO: 期待通り削除はできるけどエラーでる remove /app/images/image_01.png: no such file or directory
	if err := os.Remove(saveFileFullPath); err != nil {
		fmt.Println(err)
	}
}

func UploadToS3(localFileFullPath string, userId string) (string, string, error) {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "default",
		SharedConfigState: session.SharedConfigEnable,
	}))
	// TODO: どちらを採用するか決める
	//sess := session.Must(session.NewSession(&aws.Config{
	//	Region:      aws.String("us-east-2"),
	//	Credentials: credentials.NewSharedCredentials("/root/.aws/credentials", "default"),
	//}))

	// ローカルに保存したファイルの存在確認とポインタ取得
	file, err := os.Open(localFileFullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 仮置きファイルの拡張子から、指定するcontent-Typeを判定する
	extension := filepath.Ext(localFileFullPath)
	var contentType string
	switch extension {
	case ".jpg":
		contentType = "image/jpeg"
	case ".jpeg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".png":
		contentType = "image/png"
	default:
		return "", "", errors.New("this extension is invalid")
	}

	// オブジェクトキーの元になるuuidを作成する
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	uu := u.String()

	bucketName := "bbsapp-img"
	// ex: awsS3の {バケット}/images/123_2c2ddd5f-6571-4c8c-8f5e-b04a11250092.png
	objectKey := "images/" + userId + "_" + uu + extension

	// Uploaderを作成し、ローカルファイルをS3へアップロードする
	uploader := s3manager.NewUploader(sess)
	out, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		ContentType: aws.String(contentType),
		Body:        file,
	})
	if err != nil {
		log.Fatal(err)
	}
	objectUrl := out.Location

	return objectUrl, objectKey, err
}
