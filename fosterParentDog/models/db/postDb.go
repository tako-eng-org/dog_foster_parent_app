package db

import (
	// エンティティ(データベースのテーブルの行に対応)
	entity "fpdapp/models/entity"
	"strconv"

	// postgres用ライブラリ。importしないと下記エラーを出力する。
	// sql: unknown driver "postgres" (forgotten import?)
	_ "github.com/lib/pq"
)

//*******************************************************************
// テーブルレコードをidをキーに全件取得する
//*******************************************************************
func FindAllRecords() []entity.Post {
	allRecords := []entity.Post{}
	db := open()
	db.Order("id asc").Find(&allRecords) //select

	defer db.Close() //defer 関数がreturnする時に実行される

	return allRecords
}

//*******************************************************************
// 投稿を1ページ表示分取得する
//*******************************************************************
func FindIndexRecords(page string) []entity.Post {
	allRecords := []entity.Post{}
	db := open()
	// 数値に変換する
	pageNum, _ := strconv.Atoi(page)

	numberPerPage := 20 // 1ページあたりの表示件数

	//select
	db.Order("id desc").
		//Limit(20).Offset(0).
		Limit(numberPerPage).
		Offset((pageNum - 1) * numberPerPage).
		Find(&allRecords)

	defer db.Close() //defer 関数がreturnする時に実行される

	return allRecords
}

//*******************************************************************
// テーブルレコードをidをキーに全件取得する
//*******************************************************************
func InsertRecord(registerRecord *entity.Post) {
	db := open()
	db.Create(&registerRecord) // insert
	defer db.Close()
}
