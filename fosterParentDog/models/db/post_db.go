package db

import (
	"fpdapp/models/entity"
	"strconv"
)

//*******************************************************************
// 公開済みレコードの数を取得する
//*******************************************************************
// return(ex): -> 41
func (db *Database) CountPublishedPost() (int, error) {
	var postNum int

	//SELECT count(id) FROM "post"  WHERE (publishing = '0')
	err := db.connection.Table("posts").
		Select("count(id)").
		Where("publishing = ?", "0").
		Count(&postNum).
		Error

	return postNum, err
}

//*******************************************************************
// 公開済み投稿を1ページ表示分取得する
//*******************************************************************
func (db *Database) FindIndex(page string) ([]entity.Post, error) {
	var model []entity.Post
	pageNum, _ := strconv.Atoi(page) // 数値に変換する
	numberPerPage := 20              // 1ページあたりの表示件数

	err := db.connection.Order("id desc").
		Limit(numberPerPage).
		Offset((pageNum - 1) * numberPerPage).
		Find(&model).
		Error

	return model, err
}

//*******************************************************************
// [第1引数]の投稿IDでレコードを取得する
//*******************************************************************
func (db *Database) FindPost(postId string) (entity.Post, error) {
	var model entity.Post

	//SELECT * FROM "posts"  WHERE "posts"."deleted_at" IS NULL AND (("posts"."id" = '2')) ORDER BY "posts"."id" ASC LIMIT 1
	err := db.connection.First(&model, postId).Error

	return model, err
}

//*******************************************************************
// レコードを登録する
//*******************************************************************
func (db *Database) InsertPost(registerRecord *entity.Post) {
	db.connection.Create(&registerRecord) // insert
}
