package entity

import "github.com/jinzhu/gorm"

//投稿記事_譲渡可能都道府県テーブル用
//投稿記事に紐づく譲渡可能都道府県を格納する
type TransferablePrefecture struct {
	gorm.Model                    //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	PostId                   uint `gorm:"type:int;	not null"` //投稿記事ID
	TransferablePrefectureId int  `gorm:"type:int;	not null"` //譲渡可能都道府県ID
}
