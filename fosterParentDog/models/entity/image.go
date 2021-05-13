package entity

import "github.com/jinzhu/gorm"

// TODO: このテーブルは未完成です。後日画像処理について実装予定
// posts 1 - post_images N
// post_images 1 - images 1
type Image struct {
	gorm.Model        //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	ObjectKey  string `gorm:"type:varchar(2000)	not null"`
	UserId     uint   `gorm:"type:int;			not null"`
}
