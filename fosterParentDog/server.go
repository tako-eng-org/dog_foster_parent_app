package main

import (
	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	//// postgresql用ドライバ
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// コントローラー
	controller "fpdapp/controllers/controller"
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
	//router.Static("/dist", "./frontend/dist")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	//router.StaticFS("/bbsapp", http.Dir("./frontend/dist"))

	// 全てのJSONを返す
	router.GET("/fetchAllRecords", controller.FetchAllRecords)

	// 投稿レコード情報をDBへ登録する
	//router.POST("/addRecord", controller.AddRecord)

	router.GET("/hoge:name", func(c *gin.Context) {
		adana := c.Param("name")
		c.String(http.StatusOK, "Hello %s", adana)
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "ルートがない場合に表示するメッセージです"})
	})

	if err := router.Run(":8090"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}

}
