package main

import (
	"fpdapp/controllers/controller"
	"fpdapp/models/db"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	serve() // サーバーを起動する
}

func serve() {
	database := db.Open()
	defer database.Close()

	cont := controller.Controller{
		DbConn: database,
	}

	router := gin.Default()                   // デフォルトのミドルウェアでginのルーターを作成
	router.Static("/main", "./frontend/dist") // 静的ファイルのパスを指定

	// 公開済み投稿数を取得する  ex: localhost:8000/fosterparent/post_count?publishing=1
	router.GET("/post_count", cont.CountPost)

	// 投稿を1ページ表示分取得する ex: localhost:8000/fosterparent/index?page=1&publishing=1
	router.GET("/index", cont.IndexList)

	// 投稿を対象idの1件分取得する ex: localhost:8000/fosterparent/post?postId=1
	router.GET("/post", cont.FetchPost)

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session-id", store))

	// 投稿レコード情報をDBへ登録する
	router.POST("/post_create", cont.SessionCheck(), cont.Create)

	// 画像ファイルをS3へアップロードする
	router.POST("/image_upload", cont.SessionCheck(), cont.ImageUpload)

	router.POST("/login", cont.PostLogin)

	router.POST("/logout", cont.SessionCheck(), cont.PostLogout)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "ルートがない場合に表示するメッセージです"})
	})

	if err := router.Run(":8090"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
