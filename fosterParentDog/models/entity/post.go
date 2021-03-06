package entity

import (
	"github.com/jinzhu/gorm"
)

//投稿記事テーブル用
//ユーザーが投稿した公開及び非公開記事のテーブルである。
type Post struct {
	gorm.Model            //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	Publishing     int    `gorm:"type:int;not null"`   //公開設定
	DogName        string `gorm:"type:varchar(30);"`   //犬の名前
	Breed          string `gorm:"type:varchar(100);"`  //犬種
	Gender         int    `gorm:"type:int;"`           //性別
	Spay           int    `gorm:"type:int;"`           //去勢/避妊手術
	Old            string `gorm:"type:varchar(100);"`  //年齢
	SinglePerson   int    `gorm:"type:int;"`           //単身者への譲渡
	SeniorPerson   int    `gorm:"type:int;"`           //高齢者への譲渡
	TransferStatus int    `gorm:"type:int;"`           //譲渡ステータス
	Introduction   string `gorm:"type:varchar(2000);"` //犬の自己紹介
	AppealPoint    string `gorm:"type:varchar(2000);"` //性格アピールポイント
	OtherMessage   string `gorm:"type:varchar(4000);"` //健康状態や譲渡条件などの特記事項
	UserId         uint   `gorm:"type:int;"`           //ユーザーID

	PostImages      []PostImage      `gorm:"foreignKey:PostId"`
	PostPrefectures []PostPrefecture `gorm:"foreignKey:PostId"`
	User            User             `gorm:"foreignKey:UserId"`
}
