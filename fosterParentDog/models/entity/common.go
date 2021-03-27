package entity

import (
	"github.com/jinzhu/gorm"
)

//ユーザーテーブル
//登録ユーザーのテーブルである。
type User struct {
	gorm.Model            //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	UserName       string `gorm:"type:varchar(20);	unique	not null"	json:"user_name"` //ユーザー名
	Email          string `gorm:"type:varchar(256);	unique	not null"	json:"email"`    //メールアドレス
	Password       string `gorm:"type:varchar(128);			not null"	json:"password"`      //パスワード
	TwitterAccount string `gorm:"type:varchar(20);	unique			"	json:"twitter_account"` //twitterアカウント
	UserNickname   string `gorm:"type:varchar(30);					"	json:"user_nickname"`        //ニックネーム
	UserProfile    string `gorm:"type:varchar(400);					"	json:"user_profile"`        //自己紹介文
	WebUrl         string `gorm:"type:varchar(100);					"	json:"web_url"`             //websiteやSNSなどのURL
}

//投稿記事テーブル
//ユーザーが投稿した公開及び非公開記事のテーブルである。
type Post struct {
	gorm.Model                    //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	Publishing             int    `gorm:"type:int;			not null"	json:"publishing"`            //公開設定
	DogName                string `gorm:"type:varchar(30);			"	json:"dog_name"`              //犬の名前
	Breed                  string `gorm:"type:varchar(100);			"	json:"breed"`                //犬種
	Gender                 int    `gorm:"type:int;					"	json:"gender"`                      //性別
	Spay                   int    `gorm:"type:int;					"	json:"spay"`                        //去勢/避妊手術
	Old                    string `gorm:"type:varchar(100);			"	json:"old"`                  //年齢
	SinglePerson           int    `gorm:"type:int;					"	json:"single_person"`               //単身者への譲渡
	SeniorPerson           int    `gorm:"type:int;					"	json:"senior_person"`               //高齢者への譲渡
	TransferStatus         int    `gorm:"type:int;					"	json:"transter_status"`             //譲渡ステータス
	Introduction           string `gorm:"type:varchar(2000);		"	json:"introduction"`         //犬の自己紹介
	AppealPoint            string `gorm:"type:varchar(2000);		"	json:"appeal_point"`         //性格アピールポイント
	TransferablePrefecture int    `gorm:"type:int;					"	json:"transferable_prefecture"`     //譲渡可能都道府県ID
	OtherMessage           string `gorm:"type:varchar(4000);		"	json:"other_message"`        //健康状態や譲渡条件などの特記事項
	UserId                 uint64 `gorm:"type:int;					"	json:"user_id"`                     //ユーザーID
	TopImagePath           string `gorm:"type:varchar(150);	not null"	json:"top_image_path"` //top投稿画像パス
	PostImageId            uint64 `gorm:"type:int;					"	json:"post_image_id"`               //投稿画像ID
}

//投稿画像テーブル
//投稿記事に紐づく画像のURLを保存するテーブルである。
type PostImage struct {
	gorm.Model        //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	PostId     uint   `gorm:"type:int;			not null"	json:"post_id"`           //投稿記事ID
	ImagePath  string `gorm:"type:varchar(150);	not null"	json:"image_path"` //投稿画像パス
}

//譲渡可能都道府県テーブル
//投稿記事に紐づく譲渡可能都道府県を保存するテーブルである。
type PostPrefecture struct {
	gorm.Model        //ID, CreatedAt, UpdatedAt, DeletedAtを自動で定義する
	PostId       uint `gorm:"primary_key"	json:"post_id"`       //Post.Id
	PrefectureId uint `gorm:"primary_key"	json:"prefecture_id"` //Prefecture.Id
}

//都道府県マスターテーブル
//47都道府県一覧を保存しておくマスターテーブルである。
type Prefecture struct {
	Id   uint   `gorm:"primary_key							"	json:"id"`                  //都道府県ID
	name string `gorm:"type:varchar(10);	unique	not null"	json:"name"` //都道府県名
}

// テーブルの構造体を定義する（頭文字は大文字で始まらないとimportできないので注意）
//type FosterParentDog struct {
//	PostId     int       `gorm:"primary_key;not null"				json:"post_id"`
//	Theme      string    `gorm:"type:varchar(100);not null"		json:"theme"`
//	Title      string    `gorm:"type:varchar(200);not null"		json:"title"`
//	HandleName string    `gorm:"type:varchar(100);not null"		json:"handle_name"`
//	Detail     string    `gorm:"type:varchar(4000);not null"		json:"detail"`
//	CreatedAt  time.Time `gorm:"type:timestamp"					json:"created_at"`
//}
//type Tes struct {
//	gorm.Model
//	//キャメルケースにすると、大文字部分の前に_が自動で入る。DogName -> dog_nameのように
//	//DogName      string    `gorm:"type:varchar(100);not null"		json:"dog_name"`
//	DogNo      int    `gorm:"type:int"`
//	DogName      string    `gorm:"type:varchar(100)`
//}
