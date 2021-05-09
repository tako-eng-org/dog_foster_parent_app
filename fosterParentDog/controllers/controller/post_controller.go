package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"fpdapp/models/entity"
	"fpdapp/serializers/detail"
	"fpdapp/serializers/index"
	"fpdapp/validators/post_edit"
	"log"
	"net/http"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

//*******************************************************************
// パラメータ定義
//*******************************************************************
//公開設定
const (
	private = "0"
	public  = "1"
)

// TODO: バリデーション定数？（valudatorで実装する可能性あり）
const (
	dogNameMaxChar   int = 30
	breedMaxChar     int = 100
	genderParamRange int = 1
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

	postModel := cont.DbConn.FindPostTest(postId)
	detailSerializer := detail.Serializer{C: c, Post: postModel}

	c.JSON(http.StatusOK, detailSerializer.Response())
}

//*******************************************************************
// 投稿記事テーブルへ記事を1件登録する
//*******************************************************************
//func (cont *Controller) Create(c *gin.Context) {
//	// TODO バリデーション（投稿編集画面作成時に実装予定）
//	var post = entity.Post{
//		Publishing:     strToInt(c.PostForm("publishing")),
//		DogName:        c.PostForm("dog_name"),
//		Breed:          c.PostForm("breed"),
//		Gender:         strToInt(c.PostForm("gender")),
//		Spay:           strToInt(c.PostForm("spay")),
//		Old:            c.PostForm("old"),
//		SinglePerson:   strToInt(c.PostForm("single_person")),
//		SeniorPerson:   strToInt(c.PostForm("senior_person")),
//		TransferStatus: strToInt(c.PostForm("transfer_status")),
//		Introduction:   c.PostForm("introduction"),
//		AppealPoint:    c.PostForm("appeal_point"),
//		OtherMessage:   c.PostForm("other_message"),
//		UserId:         strToUint64(c.PostForm("user_id")),
//		TopImagePath:   c.PostForm("top_image_path"),
//	}
//	cont.DbConn.InsertPost(&post)
//	c.JSON(http.StatusCreated, post.ID)
//}
//
//*******************************************************************
// TODO: debug_投稿の新規作成
//*******************************************************************
type Post struct {
	//ID             uint   `json:"id"`
	//CreatedAt      string `json:"created_at"`
	//UpdatedAt      string `json:"updated_at"`
	Publishing     int    `json:"publishing"`      //公開設定
	DogName        string `json:"dog_name"`        //犬の名前
	Breed          string `json:"breed"`           //犬種
	Gender         int    `json:"gender"`          //性別
	Spay           int    `json:"spay"`            //去勢/避妊手術
	Old            string `json:"old"`             //年齢
	SinglePerson   int    `json:"single_person"`   //単身者への譲渡
	SeniorPerson   int    `json:"senior_person"`   //高齢者への譲渡
	TransferStatus int    `json:"transfer_status"` //譲渡ステータス
	Introduction   string `json:"introduction"`    //犬の自己紹介
	AppealPoint    string `json:"appeal_point"`    //性格アピールポイント
	OtherMessage   string `json:"other_message"`   //健康状態や譲渡条件などの特記事項
	UserId         uint64 `json:"user_id"`         //ユーザーID
	TopImagePath   string `json:"top_image_path"`  //top投稿画像パス

	ImagePathList []string `json:"image_path_list"`

	PostPrefectureIdList []int `json:"post_prefecture_id_list"`
}

func (cont *Controller) Create(c *gin.Context) {
	fmt.Println("-----------CreateTest")
	var postRequest Post
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("------BindJSONした後のPost")
	fmt.Printf("%+v\n", &postRequest)

	// postsテーブルへ登録する
	var post = entity.Post{
		Publishing:     postRequest.Publishing,
		DogName:        postRequest.DogName,
		Breed:          postRequest.Breed,
		Gender:         postRequest.Gender,
		Spay:           postRequest.Spay,
		Old:            postRequest.Old,
		SinglePerson:   postRequest.SinglePerson,
		SeniorPerson:   postRequest.SeniorPerson,
		TransferStatus: postRequest.TransferStatus,
		Introduction:   postRequest.Introduction,
		AppealPoint:    postRequest.AppealPoint,
		OtherMessage:   postRequest.OtherMessage,
		UserId:         postRequest.UserId,
		TopImagePath:   postRequest.TopImagePath,
	}
	currentPostId := cont.DbConn.InsertPost(&post)

	// post_imagesテーブルへ、パスの数分登録する
	for i := 0; i < len(postRequest.ImagePathList); i++ {
		var postImagePath = entity.PostImage{
			PostId:    currentPostId,
			ImagePath: postRequest.ImagePathList[i],
		}
		_ = cont.DbConn.InsertPostImage(&postImagePath)
	}

	// post_prefecturesテーブルへ、都道府県IDの数分登録する
	for i := 0; i < len(postRequest.PostPrefectureIdList); i++ {
		var postPrefecture = entity.PostPrefecture{
			PostId:           currentPostId,
			PostPrefectureId: postRequest.PostPrefectureIdList[i],
		}
		_ = cont.DbConn.InsertPostPrefecture(&postPrefecture)
	}

	c.JSON(http.StatusCreated, &post.ID)
}

//*******************************************************************
// TODO: debug_S3にファイルをアップロードする
//*******************************************************************
type Image struct {
	TopImagePath string `json:"top_image_path"` //top投稿画像パス
}

func (cont *Controller) ImageCreate(c *gin.Context) {
	fmt.Println("-----------ImageCreate")
	var ImageRequest Image
	if err := c.ShouldBindJSON(&ImageRequest); err != nil {
		c.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("------BindJSONした後のPost")
	fmt.Printf("%+v\n", &ImageRequest.TopImagePath)

	topImagePathBase64Pointer := &ImageRequest.TopImagePath
	topImagePathBase64 := *topImagePathBase64Pointer

	fmt.Println("------string変換した後のtopImagePathBase64")
	_, er := fmt.Println(topImagePathBase64)
	if er != nil {
		fmt.Println(er.Error())
	}

	err := UploadToS3(topImagePathBase64, "png")
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusCreated, "hogehoge")
}

type S3Service struct {
}

//func (s S3Service) UploadToS3(imageBase64 string, fileExtension string) error {
func UploadToS3(imageBase64 string, fileExtension string) error {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))

	bucketName := "bbsapp-img"
	uploader := s3manager.NewUploader(sess)

	fmt.Println("*0*0*0*0*0*0*0*0*0*0")
	fmt.Println(imageBase64)
	fmt.Println(reflect.TypeOf(imageBase64))
	data, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		fmt.Println("debug------error0")
		fmt.Println(err.Error())
	}
	wb := new(bytes.Buffer)
	wb.Write(data)

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String("foobar" + "." + fileExtension),
		Body:        wb,
		ContentType: aws.String("image/" + fileExtension),
	})
	if err != nil {
		fmt.Println(res)
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			fmt.Println("debug------error1")
			//return RaiseError(400, "Upload TimeOut", nil)
		} else {
			fmt.Println("debug------error2")
			//return RaiseError(400, "Upload Failed", nil)
		}
	}
	return nil
}

//func (cont *Controller) ImageCreate(c *gin.Context) {
//	// sessionの作成
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//		Profile:           "di",
//		SharedConfigState: session.SharedConfigEnable,
//	}))
//
//	// ファイルを開く
//	targetFilePath := "./sample.txt"
//	file, err := os.Open(targetFilePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	bucketName := "bbsapp-img" //
//	objectKey := "xxx-key"     // TODO: オブジェクトキーって何
//
//	// Uploaderを作成し、ローカルファイルをアップロード
//	uploader := s3manager.NewUploader(sess)
//	_, err = uploader.Upload(&s3manager.UploadInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String(objectKey),
//		Body:   file,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("done")
//}
