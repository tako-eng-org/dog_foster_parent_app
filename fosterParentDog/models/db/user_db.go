package db

import (
	"fpdapp/models/entity"
	"log"
)

//*******************************************************************
// [第1引数]の投稿IDで、ユーザー情報を取得する
//*******************************************************************
func (db *Database) FindUser(userId uint64) entity.User {
	var model entity.User

	err := db.connection.First(&model, userId).Error
	if err != nil {
		log.Fatal(err)
	}

	return model
}
