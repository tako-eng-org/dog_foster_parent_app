package db

import (
	"fmt"
	"fpdapp/models/entity"
	"log"
)

func (db *Database) GetUser(userName string) entity.User {
	var user entity.User
	err := db.connection.First(&user, "name = ?", userName).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}

// ユーザ数、ユーザ情報を取得
func (db *Database) UserCount(userName string) (int, entity.User) {
	var user entity.User
	cnt := 0
	err := db.connection.Where("name = ?", userName).Find(&user).Count(&cnt)
	if err != nil {
		fmt.Println(err)
	}
	return cnt, user
}

func (db *Database) UpdateSessionID(registerStruct *entity.User) entity.User {
	var user entity.User
	err := db.connection.Model(&user).
		Where("name = ?", registerStruct.Name).
		Updates(entity.User{SessionId: registerStruct.SessionId}).
		Error
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (db *Database) UpdateLastLoginAt(registerStruct *entity.User) entity.User {
	var user entity.User
	err := db.connection.Model(&user).
		Where("name = ?", registerStruct.Name).
		Updates(entity.User{LastLoginAt: registerStruct.LastLoginAt}).
		Error
	if err != nil {
		log.Fatal(err)
	}
	return user
}
