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
func (cont *Controller) CountPublishedPost(c *gin.Context) {
	count, _ := cont.DbConn.CountPublishedPost()
	c.JSON(http.StatusOK, count)
}

//*******************************************************************
// 投稿を1ページ表示(20件)分取得する
//*******************************************************************
func (cont *Controller) Index(c *gin.Context) {
	//page := c.DefaultQuery("page", "1") // ?page=1(デフォルト)
	//postModel, _ := cont.DbConn.FindIndex(page)
	//
	//fmt.Println("-------Index_s")
	//fmt.Println(postModel)
	//fmt.Println("-------Index_e")
	//
	//for i := 0; i < len(postModel); i++ {
	//	tmp := append()
	//	//PostResponce{}に、postModelの項目たちをセットする
	//}
	//
	//c.Set("my_post_model", postModel) //回避 (*conttext).MustGet: panic("Key \"" + key + "\" does not exist")
	//posts := serializers.PostSerializer{C: c, EntityPost: postModel}
	//c.JSON(http.StatusOK, posts.Response())
}

//*******************************************************************
// 投稿を対象idの1件取得する
//*******************************************************************
func (cont *Controller) FetchOnePost(c *gin.Context) {
	postId := c.Query("postId")
	postModel, _ := cont.DbConn.FindOnePost(postId) //ORMを叩いてデータとerrを取得する
	c.Set("my_post_model", postModel)               //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostSerializer{C: c, EntityPost: postModel}
	c.JSON(http.StatusOK, posts.Response())
}

//*******************************************************************
// 投稿記事テーブルへ記事を登録する
//*******************************************************************
func (cont *Controller) Create(c *gin.Context) {
	t := entity.Post{}
	var record = entity.Post{ // テーブルに登録するためのレコード情報
		Publishing:       t.ToInt(c.PostForm("publishing")),
		DogName:          t.ToStr(c.PostForm("dog_name")),
		Breed:            t.ToStr(c.PostForm("breed")),
		Gender:           t.ToInt(c.PostForm("gender")),
		Spay:             t.ToInt(c.PostForm("spay")),
		Old:              t.ToStr(c.PostForm("old")),
		SinglePerson:     t.ToInt(c.PostForm("single_person")),
		SeniorPerson:     t.ToInt(c.PostForm("senior_person")),
		TransferStatus:   t.ToInt(c.PostForm("transfer_status")),
		Introduction:     t.ToStr(c.PostForm("introduction")),
		AppealPoint:      t.ToStr(c.PostForm("appeal_point")),
		PostPrefectureId: t.ToInt(c.PostForm("transferable_prefecture")),
		OtherMessage:     t.ToStr(c.PostForm("other_message")),
		UserId:           t.ToInt64(c.PostForm("user_id")),
		TopImagePath:     t.ToStr(c.PostForm("top_image_path")),
		PostImageId:      t.ToInt64(c.PostForm("post_image_id")),
	}
	cont.DbConn.InsertPost(&record)

	c.JSON(http.StatusCreated, "****************created")
}
