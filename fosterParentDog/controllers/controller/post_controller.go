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
	page := c.DefaultQuery("page", "1") //デフォルト ?page=1
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

	postUserModel := cont.DbConn.FindPostUser(postId)
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
		Publishing:       strToInt(c.PostForm("publishing")),
		DogName:          c.PostForm("dog_name"),
		Breed:            c.PostForm("breed"),
		Gender:           strToInt(c.PostForm("gender")),
		Spay:             strToInt(c.PostForm("spay")),
		Old:              c.PostForm("old"),
		SinglePerson:     strToInt(c.PostForm("single_person")),
		SeniorPerson:     strToInt(c.PostForm("senior_person")),
		TransferStatus:   strToInt(c.PostForm("transfer_status")),
		Introduction:     c.PostForm("introduction"),
		AppealPoint:      c.PostForm("appeal_point"),
		PostPrefectureId: strToInt(c.PostForm("transferable_prefecture")),
		OtherMessage:     c.PostForm("other_message"),
		UserId:           strToInt64(c.PostForm("user_id")),
		TopImagePath:     c.PostForm("top_image_path"),
		PostImageId:      strToInt64(c.PostForm("post_image_id")),
	}
	cont.DbConn.InsertPost(&post)
	c.JSON(http.StatusCreated, post.ID)
}
