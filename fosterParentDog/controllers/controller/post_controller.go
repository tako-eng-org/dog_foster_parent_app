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
	model, _ := cont.DbConn.FindPost(c.Query("postId"))
	c.Set("my_post_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	serializer := serializers.PostSerializer{C: c, Post: model}
	c.JSON(http.StatusOK, serializer.Response())
}

//*******************************************************************
// 投稿記事テーブルへ記事を1件登録する
//*******************************************************************
func (cont *Controller) Create(c *gin.Context) {
	var post = entity.Post{
		Publishing:       ToInt(c.PostForm("publishing")),
		DogName:          ToStr(c.PostForm("dog_name")),
		Breed:            ToStr(c.PostForm("breed")),
		Gender:           ToInt(c.PostForm("gender")),
		Spay:             ToInt(c.PostForm("spay")),
		Old:              ToStr(c.PostForm("old")),
		SinglePerson:     ToInt(c.PostForm("single_person")),
		SeniorPerson:     ToInt(c.PostForm("senior_person")),
		TransferStatus:   ToInt(c.PostForm("transfer_status")),
		Introduction:     ToStr(c.PostForm("introduction")),
		AppealPoint:      ToStr(c.PostForm("appeal_point")),
		PostPrefectureId: ToInt(c.PostForm("transferable_prefecture")),
		OtherMessage:     ToStr(c.PostForm("other_message")),
		UserId:           ToInt64(c.PostForm("user_id")),
		TopImagePath:     ToStr(c.PostForm("top_image_path")),
		PostImageId:      ToInt64(c.PostForm("post_image_id")),
	}
	cont.DbConn.InsertPost(&post)
	c.JSON(http.StatusCreated, post.ID)
}
