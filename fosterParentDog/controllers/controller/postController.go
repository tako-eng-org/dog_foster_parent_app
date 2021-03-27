package controller

import (
	// 文字列と基本データ型の変換パッケージ
	//strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"
	"strconv"

	//"fpdapp/models/entity"

	// エンティティ(データベースのテーブルの行に対応)
	entity "fpdapp/models/entity"

	// DBアクセス用モジュール
	db "fpdapp/models/db"
)

//*******************************************************************
// 投稿を1ページ表示(20件)分取得する
//*******************************************************************
func FetchIndexRecords(c *gin.Context) {
	page := c.DefaultQuery("page", "1") // ?page=1(デフォルト)
	result := db.FindIndexRecords(page)
	c.JSON(200, result) // URLへのアクセスに対してJSONを返す
}

//func FetchIndexRecords(c *gin.Context) {
//  //page := c.DefaultQuery("page", "0")
//  result := db.FindIndexRecords()
//  // URLへのアクセスに対してJSONを返す
//  c.JSON(200, result)
//}

//*******************************************************************
// 投稿テーブルへ登録する
//*******************************************************************
func AddRecord(c *gin.Context) {
	reqPublishing := c.PostForm("publishing")
	reqDogName := c.PostForm("dog_name")
	reqBreed := c.PostForm("breed")
	reqGender := c.PostForm("gender")
	reqSpay := c.PostForm("spay")
	reqOld := c.PostForm("old")
	reqSinglePerson := c.PostForm("single_person")
	reqSeniorPerson := c.PostForm("senior_person")
	reqTransferStatus := c.PostForm("transter_status")
	reqIntroduction := c.PostForm("introduction")
	reqAppealPoint := c.PostForm("appeal_point")
	reqTransferablePrefecture := c.PostForm("transferable_prefecture")
	reqOtherMessage := c.PostForm("other_message")
	reqUserId := c.PostForm("user_id")
	reqTopImagePath := c.PostForm("top_image_path")
	reqPostImageId := c.PostForm("post_image_id")

	// 数値に変換する
	publishing, _ := strconv.Atoi(reqPublishing)
	gender, _ := strconv.Atoi(reqGender)
	spay, _ := strconv.Atoi(reqSpay)
	singlePerson, _ := strconv.Atoi(reqSinglePerson)
	seniorPerson, _ := strconv.Atoi(reqSeniorPerson)
	transferStatus, _ := strconv.Atoi(reqTransferStatus)
	transfablePrefecture, _ := strconv.Atoi(reqTransferablePrefecture)
	userId, _ := strconv.ParseUint(reqUserId, 10, 64)
	postImageId, _ := strconv.ParseUint(reqPostImageId, 10, 64)

	// テーブルに登録するためのレコード情報
	var record = entity.Post{
		Publishing:             publishing,
		DogName:                reqDogName,
		Breed:                  reqBreed,
		Gender:                 gender,
		Spay:                   spay,
		Old:                    reqOld,
		SinglePerson:           singlePerson,
		SeniorPerson:           seniorPerson,
		TransferStatus:         transferStatus,
		Introduction:           reqIntroduction,
		AppealPoint:            reqAppealPoint,
		TransferablePrefecture: transfablePrefecture,
		OtherMessage:           reqOtherMessage,
		UserId:                 userId,
		TopImagePath:           reqTopImagePath,
		PostImageId:            postImageId,
	}
	db.InsertRecord(&record)
}
