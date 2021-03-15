package controller

import (
	// 文字列と基本データ型の変換パッケージ
	//strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"
	//"fpdapp/models/entity"

	// エンティティ(データベースのテーブルの行に対応)
	//entity "../../models/entity"

	// DBアクセス用モジュール
	db "fpdapp/models/db"
)

// 全ての投稿情報を取得する
func FetchAllRecords(c *gin.Context) {
	result := db.FindAllRecords()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}
