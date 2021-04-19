package db

import (
	"fpdapp/models/entity"
	"log"
)

//*******************************************************************
// [第1引数]の投稿IDで、譲渡可能都道府県を取得する
//*******************************************************************
func (db *Database) FindPostPrefectures(postId string) []entity.PostPrefecture {
	var model []entity.PostPrefecture

	err := db.connection.Where("post_id = ?", postId).Find(&model).Error
	if err != nil {
		log.Fatal(err)
	}

	return model
}
