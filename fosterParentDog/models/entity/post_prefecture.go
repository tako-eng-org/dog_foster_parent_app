package entity

import "github.com/jinzhu/gorm"

//譲渡可能都道府県テーブル用
//投稿記事に紐づく譲渡可能都道府県を保存するテーブルである。
type PostPrefecture struct {
	gorm.Model        //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	PostId       uint `gorm:"primary_key"` //Post.Id
	PrefectureId uint `gorm:"primary_key"` //Prefecture.Id
}
