package entity

import "github.com/jinzhu/gorm"

// PostとImageの中間テーブル
type PostImage struct {
	gorm.Model      //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	PostId     uint `gorm:"type:int; not null"`
	Position   int  `gorm:"type:int; default:0 not null"`

	// FIXME: 少し書いているけど、この定義はまだdb処理を追加していません。後日改修予定
	ImageId uint    `gorm:"type:int;"`
	Images  []Image `gorm:"foreignKey:ImageId"`
}
