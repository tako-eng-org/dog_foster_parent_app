package db

import (
	entity "fpdapp/models/entity"
	"strconv"

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
// 公開済みレコードの数を取得する
//*******************************************************************
// return(ex): -> 41
func CountPublishedPostNum() int64 {
	table := []entity.Post{}
	db := open()

	var count int64

	//SELECT count(id) FROM "post"  WHERE "post"."deleted_at" IS NULL AND ((publishing = '0'))
	db.Model(&table).
		Select("count(id)").
		Where("publishing = ?", "0").
		Count(&count)

	defer db.Close() //defer 関数がreturnする時に実行される

	return count
}

//*******************************************************************
// 投稿を1ページ表示分取得する
//*******************************************************************
func FindIndexRecords(page string) []entity.Post {
	records := []entity.Post{}
	db := open()
	// 数値に変換する
	pageNum, _ := strconv.Atoi(page)

	numberPerPage := 20 // 1ページあたりの表示件数

	//SELECT * FROM "post"  WHERE "post"."deleted_at" IS NULL ORDER BY id desc LIMIT 20 OFFSET 0
	db.Order("id desc").
		Limit(numberPerPage).
		Offset((pageNum - 1) * numberPerPage).
		Find(&records)

	defer db.Close() //defer 関数がreturnする時に実行される

	return records
}

//*******************************************************************
//
//*******************************************************************
func InsertRecord(registerRecord *entity.Post) {
	db := open()
	db.Create(&registerRecord) // insert
	defer db.Close()
}
