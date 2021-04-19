package db

import (
	"fpdapp/models/entity"
	"log"
)

//*******************************************************************
// [第1引数]の投稿IDで、投稿画像テーブルから投稿画像パスを取得する
//*******************************************************************
func (db *Database) FindPostImages(postId string) []entity.PostImage {
	var model []entity.PostImage

	err := db.connection.Where("post_id = ?", postId).Find(&model).Error
	if err != nil {
		log.Fatal(err)
	}

	return model
}
