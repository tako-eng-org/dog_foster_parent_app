package entity

//都道府県マスターテーブル用
//47都道府県一覧を保存しておくマスターテーブルである。
type Prefecture struct {
	Id   uint   `gorm:"primary_key							"`                //都道府県ID
	Name string `gorm:"type:varchar(10);	unique	not null"` //都道府県名
	Kana string `gorm:"type:varchar(20);	unique	not null"` //都道府県読み仮名
}
