package main

import (
	"fpdapp/controllers/controller"
	"fpdapp/models/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	serve() // サーバーを起動する
}

func serve() {
	database := db.Open()
	defer database.Close()

	postController := controller.PostController{
		Database: database,
	}

	router := gin.Default()                   // デフォルトのミドルウェアでginのルーターを作成
	router.Static("/main", "./frontend/dist") // 静的ファイルのパスを指定

	// ************************************************
	// トップ画面から使用するAPI
	// ************************************************
	// 公開済み投稿数を取得する
	router.GET("/published_post_count", postController.CountPublishedPost)

	// 投稿を1ページ表示分取得する
	// ex: localhost:8000/fosterparent/index?page=1
	router.GET("/index", postController.Index)

	// ************************************************
	// 投稿詳細画面から使用するAPI
	// ************************************************
	// 投稿を対象idの1件分取得する
	// ex: localhost:8000/fosterparent/post?postId=1
	router.GET("/post", postController.FetchOnePost)

	// 投稿idをもとに、投稿画像を取得する
	// ex: localhost:8000/fosterparent/images?postId=44
	router.GET("/images", postController.FetchPostImagePaths)

	// 投稿idをもとに、投稿画像を取得する
	// ex: localhost:8000/fosterparent/detail/44
	router.GET("/post_prefecture", postController.FetchPostPrefecture)

	// 投稿idをもとに、投稿画像を取得する
	// ex: localhost:8000/fosterparent/?????????
	router.GET("/user", postController.FetchPostUser)

	// ************************************************
	// test
	// ************************************************
	// 投稿レコード情報をDBへ登録する
	router.POST("/add_record", postController.Create)

	// ************************************************
	// 異常系
	// ************************************************
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "ルートがない場合に表示するメッセージです"})
	})

	if err := router.Run(":8090"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
