package entity

import "github.com/jinzhu/gorm"

//投稿画像テーブル用
//投稿記事に紐づく画像のURLを保存するテーブルである。
type PostImage struct {
	gorm.Model        //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	PostId     uint   `gorm:"type:int;			not null"`        //投稿記事ID
	ImagePath  string `gorm:"type:varchar(150);	not null"` //投稿画像パス
}
