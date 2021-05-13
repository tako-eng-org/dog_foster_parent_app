package entity

import "github.com/jinzhu/gorm"

//ユーザーテーブル用
//登録ユーザーのテーブルである。
type User struct {
	gorm.Model            //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	Name           string `gorm:"type:varchar(20); unique not null"`  //ユーザー名
	Email          string `gorm:"type:varchar(256);	unique not null"` //メールアドレス
	Password       string `gorm:"type:varchar(128);	not null"`        //パスワード
	TwitterAccount string `gorm:"type:varchar(20); unique"`           //twitterアカウント
	Nickname       string `gorm:"type:varchar(30);"`                  //ニックネーム
	Profile        string `gorm:"type:varchar(400);"`                 //自己紹介文
	WebUrl         string `gorm:"type:varchar(100);"`                 //websiteやSNSなどのURL
}
