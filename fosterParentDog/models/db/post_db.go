package db

import (
	"fpdapp/models/entity"
	"log"
	"strconv"
)

//*******************************************************************
// 公開済みレコードの数を取得する ex:return -> 41
//*******************************************************************
func (db *Database) CountPost(publishing string) int {
	var postNum int

	//SELECT count(id) FROM "post"  WHERE (publishing = '0')
	err := db.connection.Table("posts").
		Select("count(id)").
		Where("publishing = ?", publishing).
		Count(&postNum).
		Error
	if err != nil {
		log.Fatal(err)
	}

	return postNum
}

//*******************************************************************
// 公開済み投稿を1ページ表示分取得する
//*******************************************************************
func (db *Database) FindIndex(page string, publishing string) []entity.Post {
	pageNum, _ := strconv.Atoi(page)
	numberPerPage := 20 // 1ページあたりの表示件数

	var result []entity.Post

	err := db.connection.
		Order("id desc").
		Where("publishing = ?", publishing).
		Preload("PostImages").
		Preload("PostPrefectures").
		Preload("User").
		Limit(numberPerPage).
		Offset((pageNum - 1) * numberPerPage).
		Find(&result).
		Error
	if err != nil {
		log.Fatal(err)
	}

	return result
}

//*******************************************************************
// [第1引数]の投稿IDでレコードを取得する
//*******************************************************************
func (db *Database) FindPost(postId string) entity.Post {
	var model entity.Post

	//SELECT * FROM "posts"  WHERE "posts"."deleted_at" IS NULL AND (("posts"."id" = '2')) ORDER BY "posts"."id" ASC LIMIT 1
	err := db.connection.
		Preload("PostImages").
		Preload("PostPrefectures").
		Preload("User").
		First(&model, postId).
		Error
	if err != nil {
		log.Fatal(err)
	}

	return model
}

//*******************************************************************
// レコードを登録する
//*******************************************************************
func (db *Database) InsertPost(registerStruct *entity.Post) uint {
	err := db.connection.
		Create(&registerStruct). // preload不要
		Error
	if err != nil {
		log.Fatal(err)
	}
	return registerStruct.ID //登録後レコードのID(post.ID)
}

//*******************************************************************
// 画像情報を登録する
//*******************************************************************
// TODO: この関数は機能するが、画像テーブルの定義を見直し後に改修する必要がある。
func (db *Database) InsertPostImage(targetStruct *entity.PostImage) *entity.PostImage {
	err := db.connection.Create(&targetStruct).Error
	if err != nil {
		log.Fatal(err)
	}
	return targetStruct
}
