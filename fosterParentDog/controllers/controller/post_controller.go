package controller

import (
	db "fpdapp/models/db"
	"fpdapp/models/entity"
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	PostController struct {
		Database *db.Database
	}
)

//*******************************************************************
// 公開済み投稿数を取得する
//*******************************************************************
func (pc *PostController) CountPublishedPostNum(c *gin.Context) {
	c.JSON(http.StatusOK, pc.Database.CountPublishedPostNum()) // URLへのアクセスに対してJSONを返す
}

//*******************************************************************
// 投稿を1ページ表示(20件)分取得する
//*******************************************************************
func (pc *PostController) Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1") // ?page=1(デフォルト)
	postModel, _ := pc.Database.FindIndexRecords(page)
	c.Set("my_post_model", postModel) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostSerializer{C: c, EntityPost: postModel}
	c.JSON(http.StatusOK, posts.Response())
}

//*******************************************************************
// 投稿テーブルへ登録する
//*******************************************************************
func (pc *PostController) Create(c *gin.Context) {
	t := entity.Post{}
	var record = entity.Post{ // テーブルに登録するためのレコード情報
		Publishing:             t.SetPublishing(c.PostForm("publishing")),
		DogName:                t.SetDogName(c.PostForm("dog_name")),
		Breed:                  t.SetBreed(c.PostForm("breed")),
		Gender:                 t.SetGender(c.PostForm("gender")),
		Spay:                   t.SetSpay(c.PostForm("spay")),
		Old:                    t.SetOld(c.PostForm("old")),
		SinglePerson:           t.SetSinglePerson(c.PostForm("single_person")),
		SeniorPerson:           t.SetSeniorPerson(c.PostForm("senior_person")),
		TransferStatus:         t.SetTransferablePrefecture(c.PostForm("transter_status")),
		Introduction:           t.SetIntroduction(c.PostForm("introduction")),
		AppealPoint:            t.SetAppealPoint(c.PostForm("appeal_point")),
		TransferablePrefecture: t.SetTransferablePrefecture(c.PostForm("transferable_prefecture")),
		OtherMessage:           t.SetOtherMessage(c.PostForm("other_message")),
		UserId:                 t.SetUserId(c.PostForm("user_id")),
		TopImagePath:           t.SetTopImagePath(c.PostForm("top_image_path")),
		PostImageId:            t.SetPostImageId(c.PostForm("post_image_id")),
	}
	pc.Database.InsertRecord(&record)

	c.JSON(http.StatusCreated, "****************created")
}

//*******************************************************************
// 投稿を対象idの1件取得する
//*******************************************************************
func (pc *PostController) FetchPostOneRecord(c *gin.Context) {
	postId := c.Query("postId")
	postModel, _ := pc.Database.FindPostOneRecord(postId)            //ORMを叩いてデータとerrを取得する
	c.Set("my_post_model", postModel)                                //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostSerializer{C: c, EntityPost: postModel} //結果をコンテキストとスライス[]に格納する
	c.JSON(http.StatusOK, posts.Response())                          //コンテキストに入ったデータを整形してリターン
}

//*******************************************************************
// 投稿idをもとに、投稿画像を取得する
//*******************************************************************
func (pc *PostController) FetchPostImagePaths(c *gin.Context) {
	postId := c.Query("postId")
	postImageModel, _ := pc.Database.FindPostImagePaths(postId) //ORMを叩いてデータとerrを取得する

	c.Set("my_post_image_model", postImageModel)                                    //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostImageSerializer{C: c, EntityPostImage: postImageModel} //結果をコンテキストとスライス[]に格納する
	c.JSON(http.StatusOK, posts.ResponsePostImage())                                //コンテキストに入ったデータを整形してリターン
}

//*******************************************************************
// 投稿idをもとに、投稿画像を取得する
//*******************************************************************
func (pc *PostController) FetchPostTransferablePrefecture(c *gin.Context) {
	postId := c.Query("postId")
	model, _ := pc.Database.FindPostFetchPostTransferablePrefecture(postId) //ORMを叩いてデータとerrを取得する

	c.Set("my_post_transferable_prefecture_model", model)                                                    //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostTransferablePrefectureSerializer{C: c, EntityPostTransferablePrefecture: model} //結果をコンテキストとスライス[]に格納する
	c.JSON(http.StatusOK, posts.ResponsePostTransferablePrefecture())                                        //コンテキストに入ったデータを整形してリターン
}

//*******************************************************************
// 投稿idをもとに、user情報を取得する
//*******************************************************************
func (pc *PostController) FetchPostUserProfile(c *gin.Context) {
	postId := c.Query("postId")
	model, _ := pc.Database.FindPostUserProfile(postId) //ORMを叩いてデータとerrを取得する

	c.Set("my_post_user_model", model)            //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.UserSerializer{c, model} //結果をコンテキストとスライス[]に格納する
	c.JSON(http.StatusOK, posts.ResponseUser())   //コンテキストに入ったデータを整形してリターン
}
