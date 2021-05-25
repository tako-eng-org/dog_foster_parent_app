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

// TODO: session-idがない前提のログイン処理
func (cont *Controller) PostLogin(c *gin.Context) {
	//cookieSessionId, _ := c.Cookie("session-id")
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

	// ログイン成功した場合は、セッションIDとセッションユーザ名を発行
	session := sessions.Default(c)
	//dbUser.SessionId = cookieSessionId
	//dbUser.SessionId = session.Get("session-id")
	//log.Println("--------------------")
	//log.Println(session.Get("session-id"))
	//log.Println("--------------------")
	//dbUser = cont.DbConn.UpdateSessionID(&dbUser)
	c.SetCookie("user-name", dbUser.Name, 3600, "/", "localhost", false, true) //これがないとcookieにuser-nameがセットされない。
	session.Set("user-name", dbUser.Name)                                      //これがないとsession-idがセットされない。
	err := session.Save()
	if err != nil {
		log.Println("セッションの保存に失敗しました：", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "予期せぬエラー。セッションの保存に失敗しました"})
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{"user name is:": session.Get("user-name")})
}

func (cont *Controller) SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		log.Println("---session is :", session)
		cookieUserName, _ := c.Cookie("user-name")
		cookieSessionId, _ := c.Cookie("session-id")
		log.Println("-----------------")
		log.Println("Cookieのuser-nameは:", cookieUserName)
		log.Println("Cookieのsession-idは:", cookieSessionId)
		log.Println("sessionのユーザ名は:", session.Get("user-name"))
		log.Println("-----------------")
		var dbUser entity.User
		_, dbUser = cont.DbConn.UserCount(cookieUserName)

		//if dbUser.SessionId == cookieSessionId &&
		//	dbUser.Name == cookieUserName {
		if dbUser.Name == session.Get("user-name") {
			// セッション一致。ログイン不要
			log.Println("DBのセッションID&ユーザ名と、cookieの情報が一致しました。すでにログイン済です。")
			//c.Set("session-id", cookieSessionId)
			//c.Set("user-name", cookieUserName)
			// ユーザーの最終ログイン日時を更新する
			dbUser = entity.User{Name: cookieUserName, LastLoginAt: time.Now()}
			cont.DbConn.UpdateLastLoginAt(&dbUser)
		} else {
			// セッションなし。ログインが必要
			log.Println("DBのセッションID&ユーザ名と、cookieの情報が一致しません。")
			c.JSON(http.StatusBadRequest, gin.H{"error": "ログインしていません。"})
			c.Abort()
		}
	}
}

func (cont *Controller) PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	//セッションからデータを破棄する
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		log.Println("セッションの破棄に失敗しました：", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "予期せぬエラー。セッションの破棄に失敗しました"})
		c.Abort()
	}

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
