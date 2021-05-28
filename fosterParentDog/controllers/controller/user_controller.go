package controller

import (
	"fpdapp/models/entity"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (cont *Controller) GetUserMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"テスト": "ユーザーログイン後の画面"})
}

func (cont *Controller) PostLogin(c *gin.Context) {
	formUserName := c.PostForm("user-name")
	formPassword := c.PostForm("password")
	var dbUser entity.User
	userCount, dbUser := cont.DbConn.UserCount(formUserName)

	// 注意：ユーザ名がDBに存在しない場合&パスワード不一致の場合、同じエラーメッセージで返すこと。（ブルートフォース攻撃対策）
	// フォームユーザ名がDBに存在しない場合、ログインさせない
	if userCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "ユーザかパスワードが存在しません。登録してください。"})
		c.Abort()
	}

	// formパスワードとDBパスワードが不一致の場合、ログインさせない
	if err := compareHashAndPassword(dbUser.Password, formPassword); err != nil {
		log.Println("パスワード不一致：", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "ユーザかパスワードが存在しません。登録してください。"})
		c.Abort()
	}

	// ログイン成功した場合は、セッション内にユーザ名を発行＆クッキーにユーザ名をセットする
	session := sessions.Default(c)
	session.Set("user-name", dbUser.Name) //これがないとsession-idがセットされない。
	err := session.Save()
	if err != nil {
		log.Println("セッションの保存に失敗しました：", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "予期せぬエラー。セッションの保存に失敗しました"})
		c.Abort()
	}
	c.SetCookie("user-name", dbUser.Name, 3600, "/", "localhost", false, true) //これがないとcookieにuser-nameがセットされない。

	// ユーザーの最終ログイン日時を更新する
	dbUser = entity.User{Name: formUserName, LastLoginAt: time.Now()}
	cont.DbConn.UpdateLastLoginAt(&dbUser)

	c.JSON(http.StatusOK, gin.H{"user name is:": session.Get("user-name")})
}

func (cont *Controller) SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		cookieUserName, _ := c.Cookie("user-name")
		var dbUser entity.User
		_, dbUser = cont.DbConn.UserCount(cookieUserName)

		// セッションに入っているユーザ名と、DBのユーザ名が一致していればログイン済
		if dbUser.Name == session.Get("user-name") {
			// セッション一致。ログイン済
			log.Println("すでにログイン済です。")
			// ユーザーの最終ログイン日時を更新する
			dbUser = entity.User{Name: cookieUserName, LastLoginAt: time.Now()}
			cont.DbConn.UpdateLastLoginAt(&dbUser)
		} else {
			// セッションなし。ログインが必要
			log.Println("セッションのユーザ名と、DBのユーザ名が一致しません。")
			c.JSON(http.StatusBadRequest, gin.H{"error": "ログインしていません。"})
			c.Abort()
		}
		log.Println("セッションチェック終了")
	}
}

func (cont *Controller) PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	//セッションからデータを破棄する
	session := sessions.Default(c)
	session.Set("user-name", "") //session-idのsetはしなくても削除される。
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	err := session.Save()
	if err != nil {
		log.Println("セッションの破棄に失敗しました：", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "予期せぬエラー。セッションの破棄に失敗しました"})
		c.Abort()
	}
	c.SetCookie("user-name", "", -1, "/", "localhost", false, true) // maxAge:-1 はクッキーの削除を表す
	c.JSON(http.StatusOK, gin.H{"": "ログアウトしました"})
}

// パスワードをhash化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// hashと非hashパスワード比較
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
