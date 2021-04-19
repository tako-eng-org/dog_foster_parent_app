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

	//select post_prefecture.id, post_prefecture.post_prefecture_id from post_prefecture where post_id = 44;
	err := db.connection.Table("post_prefectures").
		Select("post_prefectures.id,"+
			" post_prefectures.post_id,"+
			" post_prefectures.post_prefecture_id").
		Where("post_prefectures.post_id = ?", postId).
		Scan(&model).
		Error
	if err != nil {
		log.Fatal(err)
	}

	return model
}
