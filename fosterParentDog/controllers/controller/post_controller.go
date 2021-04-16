package controller

import (
	"fpdapp/models/entity"
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 公開済み投稿数を取得する
//*******************************************************************
// TODO 今更思ったけど、 published_post_count のエンドポイントじゃなくて post_count のエンドポイントを作成してクエリで公開、非公開を指定できた方が汎用性高くないかな？
// https://github.com/tako-eng-org/dog_foster_parent_app/pull/24#discussion_r611191794
func (cont *Controller) CountPublishedPost(c *gin.Context) {
	count, _ := cont.DbConn.CountPublishedPost()
	c.JSON(http.StatusOK, count)
}

//*******************************************************************
// 投稿を1ページ表示件数分取得する
//*******************************************************************
func (cont *Controller) IndexList(c *gin.Context) {
	page := c.DefaultQuery("page", "1") // ?page=1(デフォルト)
	model, _ := cont.DbConn.FindIndex(page)
	c.Set("my_post_prefecture_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	serializer := serializers.PostsSerializer{C: c, Posts: model}
	c.JSON(http.StatusOK, serializer.Response())
}

//*******************************************************************
// 投稿を対象idの1件取得する
//*******************************************************************
func (cont *Controller) FetchPost(c *gin.Context) {
	postId := c.Query("postId")

	// post
	postModel, _ := cont.DbConn.FindPost(postId)
	postSerializer := serializers.PostSerializer{C: c, Post: postModel}

	// postImage
	postImageModel, _ := cont.DbConn.FindPostImagePathList(postId)
	postImageSerializer := serializers.PostImagesSerializer{C: c, PostImages: postImageModel}

	// postPrefecture
	postPrefectureListModel, _ := cont.DbConn.FindPostPrefectureList(postId)
	postPrefectureListSerializer := serializers.PostPrefecturesSerializer{C: c, PostPrefectures: postPrefectureListModel}

	// postUser
	postUserModel, _ := cont.DbConn.FindPostUser(postId)
	postUserSerializer := serializers.UserSerializer{C: c, User: postUserModel}

	type Response struct {
		Post            serializers.PostResponse             `json:"post"`
		PostImages      []serializers.PostImageResponse      `json:"post_images"`
		PostPrefectures []serializers.PostPrefectureResponse `json:"post_prefectures"`
		User            serializers.UserResponse             `json:"user"`
	}

	var r Response
	r.Post = postSerializer.Response()
	r.PostImages = postImageSerializer.Response()
	r.PostPrefectures = postPrefectureListSerializer.Response()
	r.User = postUserSerializer.Response()

	c.JSON(http.StatusOK, r)
}

//*******************************************************************
// 投稿記事テーブルへ記事を1件登録する
//*******************************************************************
func (cont *Controller) Create(c *gin.Context) {
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
