package db

import (
	"fmt"
	"fpdapp/models/entity"
)

//*******************************************************************
// [第1引数]の投稿IDで、ユーザー情報を取得する
//*******************************************************************
func (db *Database) FindPostUser(postId string) (entity.User, error) {
	var model entity.User

	/*
		select B.*
		from post A
		inner join public.user B
		on A.user_id = B.id
		where A.id = 44;
	*/
	err := db.connection.Table("posts A").
		Select("B.*").
		Joins("inner join public.users B on A.user_id = B.id").
		Where("A.id = ?", postId).
		Scan(&model).
		Error

	if err != nil {
		fmt.Println(err)
	}

	return model, err
}
