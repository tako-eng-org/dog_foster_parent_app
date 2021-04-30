package controller

import (
	"fpdapp/models/entity"
	"fpdapp/serializers"
	"net/http"

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
	model := cont.DbConn.FindIndex(page, publishing)
	serializer := serializers.PostsSerializer{C: c, Posts: model}
	c.JSON(http.StatusOK, serializer.Response())
}

//*******************************************************************
// 投稿を対象idの1件取得する
//*******************************************************************
func (cont *Controller) FetchPost(c *gin.Context) {
	postId := c.Query("postId")

	postModel := cont.DbConn.FindPost(postId)
	postSerializer := serializers.PostSerializer{C: c, Post: postModel}

	postImageModel := cont.DbConn.FindPostImages(postId)
	postImageSerializer := serializers.PostImagesSerializer{C: c, PostImages: postImageModel}

	postPrefectureListModel := cont.DbConn.FindPostPrefectures(postId)
	postPrefectureListSerializer := serializers.PostPrefecturesSerializer{C: c, PostPrefectures: postPrefectureListModel}

	postUserModel := cont.DbConn.FindUser(postModel.UserId)
	postUserSerializer := serializers.UserSerializer{C: c, User: postUserModel}

	response := struct {
		Post            serializers.PostResponse             `json:"post"`
		PostImages      []serializers.PostImageResponse      `json:"post_images"`
		PostPrefectures []serializers.PostPrefectureResponse `json:"post_prefectures"`
		User            serializers.UserResponse             `json:"user"`
	}{
		postSerializer.Response(),
		postImageSerializer.Response(),
		postPrefectureListSerializer.Response(),
		postUserSerializer.Response(),
	}

	c.JSON(http.StatusOK, response)
}

//*******************************************************************
// 投稿記事テーブルへ記事を1件登録する
//*******************************************************************
func (cont *Controller) Create(c *gin.Context) {
	// TODO バリデーション（投稿編集画面作成時に実装予定）
	var post = entity.Post{
		Publishing:     strToInt(c.PostForm("publishing")),
		DogName:        c.PostForm("dog_name"),
		Breed:          c.PostForm("breed"),
		Gender:         strToInt(c.PostForm("gender")),
		Spay:           strToInt(c.PostForm("spay")),
		Old:            c.PostForm("old"),
		SinglePerson:   strToInt(c.PostForm("single_person")),
		SeniorPerson:   strToInt(c.PostForm("senior_person")),
		TransferStatus: strToInt(c.PostForm("transfer_status")),
		Introduction:   c.PostForm("introduction"),
		AppealPoint:    c.PostForm("appeal_point"),
		OtherMessage:   c.PostForm("other_message"),
		UserId:         strToUint64(c.PostForm("user_id")),
		TopImagePath:   c.PostForm("top_image_path"),
	}
	cont.DbConn.InsertPost(&post)
	c.JSON(http.StatusCreated, post.ID)
}

//*******************************************************************
// TODO: debug_投稿の新規作成
//*******************************************************************
//type PostBase struct {
//	//ID             uint   `json:"id"`
//	//CreatedAt      string `json:"created_at"`
//	//UpdatedAt      string `json:"updated_at"`
//	Publishing     int    `json:"publishing"`      //公開設定
//	DogName        string `json:"dog_name"`        //犬の名前
//	Breed          string `json:"breed"`           //犬種
//	Gender         int    `json:"gender"`          //性別
//	Spay           int    `json:"spay"`            //去勢/避妊手術
//	Old            string `json:"old"`             //年齢
//	SinglePerson   int    `json:"single_person"`   //単身者への譲渡
//	SeniorPerson   int    `json:"senior_person"`   //高齢者への譲渡
//	TransferStatus int    `json:"transfer_status"` //譲渡ステータス
//	Introduction   string `json:"introduction"`    //犬の自己紹介
//	AppealPoint    string `json:"appeal_point"`    //性格アピールポイント
//	OtherMessage   string `json:"other_message"`   //健康状態や譲渡条件などの特記事項
//	UserId         uint64 `json:"user_id"`         //ユーザーID
//	TopImagePath   string `json:"top_image_path"`  //top投稿画像パス
//}
//type PostImage struct {
//	//PostID      uint   `json:"post_id"`
//	//PostImageID uint   `json:"post_image_id"`
//	ImagePath string `json:"image_path"`
//}
//type PostPrefecture struct {
//	//PostId           uint `json:"post_id"`
//	PostPrefectureId int `json:"post_prefecture_id"`
//}
//
//type StTest struct {
//	PostBase        `json:"postBase"`
//	PostImage          `json:"imagePathList"`
//	PostPrefecture  `json:"postPrefectureList"`
//}

type St struct {
	Post           serializers.PostResponse           `json:"postBase"`
	PostImage      serializers.PostImageResponse      `json:"imagePathList"`
	PostPrefecture serializers.PostPrefectureResponse `json:"postPrefectureList"`
}

func (cont *Controller) CreateTest(c *gin.Context) {
	fmt.Println("-----------CreateTest")
	fmt.Println(c)
	//var postTest entity.Post
	var postStruct St
	c.BindJSON(&postStruct)

	fmt.Println("------BindJSONした後のPost")
	fmt.Printf("%+v", &postStruct)

	c.JSON(http.StatusCreated, "----status----")
	//cont.DbConn.InsertPostTest(&postTest)

}

//*******************************************************************
// TODO: debug_S3にファイルをアップロードする
//*******************************************************************
//func (cont *Controller) postImageTest(c *gin.Context) {
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
//	bucketName := "bbsapp-img"
//	objectKey := "xxx-key" // TODO: オブジェクトキーって何
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
