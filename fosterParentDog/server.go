package main

import (
	"fpdapp/controllers/controller"
	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	//// postgresql用ドライバ
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// 静的ファイルのパスを指定
	router.Static("/dist", "./frontend/dist")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	router.StaticFS("/main", http.Dir("./frontend/dist"))

	// ************************************************
	// 共通（デバッグ用）
	// ************************************************
	router.GET("/fetchAllRecords", controller.FetchAllRecords)

	// ************************************************
	// トップ画面から使用するAPI
	// ************************************************
	// 公開済み投稿数を取得する
	router.GET("/pageCount", controller.FetchPublishedPostNum)

	// 投稿を1ページ表示分取得する
	// ex: localhost:8000/fosterparent/index?page=1
	router.GET("/index", controller.FetchIndexRecords)

	// ************************************************
	// 投稿詳細画面から使用するAPI
	// ************************************************
	// 投稿レコード情報をDBへ登録する
	router.POST("/addRecord", controller.AddRecord)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "ルートがない場合に表示するメッセージです"})
	})

	if err := router.Run(":8090"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
