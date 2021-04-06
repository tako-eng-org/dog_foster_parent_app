package controller

import (
	"fmt"
	db "fpdapp/models/db"
	entity "fpdapp/models/entity"
	"fpdapp/serializers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 公開済み投稿数を取得する
//*******************************************************************
func CountPublishedPostNum(c *gin.Context) {
	c.JSON(200, db.CountPublishedPostNum()) // URLへのアクセスに対してJSONを返す
}

//*******************************************************************
// 投稿を1ページ表示(20件)分取得する
//*******************************************************************
func Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")               // ?page=1(デフォルト)
	postModel, _ := db.FindIndexRecords(page)         //ORMを叩いてデータとerrを取得する
	c.Set("my_post_model", postModel)                 //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostSerializer{c, postModel} //結果をコンテキストとスライス[]に格納する
	c.JSON(http.StatusOK, posts.Response())           //コンテキストに入ったデータを整形してリターン
}

//*******************************************************************
// 投稿テーブルへ登録する
//*******************************************************************
func Create(c *gin.Context) {
	userId, _ := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
	postImageId, _ := strconv.ParseUint(c.PostForm("post_image_id"), 10, 64)

	// テーブルに登録するためのレコード情報
	var record = entity.Post{
		Publishing:             strToInt(c.PostForm("publishing")),
		DogName:                c.PostForm("dog_name"),
		Breed:                  c.PostForm("breed"),
		Gender:                 strToInt(c.PostForm("gender")),
		Spay:                   strToInt(c.PostForm("spay")),
		Old:                    c.PostForm("old"),
		SinglePerson:           strToInt(c.PostForm("single_person")),
		SeniorPerson:           strToInt(c.PostForm("senior_person")),
		TransferStatus:         strToInt(c.PostForm("transter_status")),
		Introduction:           c.PostForm("introduction"),
		AppealPoint:            c.PostForm("appeal_point"),
		TransferablePrefecture: strToInt(c.PostForm("transferable_prefecture")),
		OtherMessage:           c.PostForm("other_message"),
		UserId:                 userId,
		TopImagePath:           c.PostForm("top_image_path"),
		PostImageId:            postImageId,
	}
	db.InsertRecord(&record)
}

// 文字列を数値に変換する
func strToInt(arg string) int {
	ret, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("This function was error: %q\n", err)
	}
	return ret
}
