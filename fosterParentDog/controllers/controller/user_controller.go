package controller

import (
	"fmt"
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
	log.Println("------------main logic")
	cookieUserName, _ := c.Cookie("user_name")
	cookiePassword, _ := c.Cookie("password")
	cookieSessionId, _ := c.Cookie("session-id")
	//formUserName := c.PostForm("user_name")
	//formPassword := c.PostForm("password")
	var dbUser entity.User
	userCount, dbUser := cont.DbConn.UserCount(cookieUserName)
	fmt.Println("--------usercount ", userCount)

	// フォームユーザ名がDBに存在しない場合、ログインさせない
	if userCount == 0 { // ユーザ名がDBに存在しない場合、エラー
		c.JSON(http.StatusBadRequest, gin.H{"err": "ユーザが存在しません。登録してください。"})
		c.Abort()
	} else { // ユーザ名が存在する場合、次の処理へ
		log.Println("ユーザが存在します。ログイン処理を続行します")
	}

	// formパスワードとDBパスワードが一致した場合、ログイン処理継続
	if err := compareHashAndPassword(dbUser.Password, cookiePassword); err != nil {
		log.Println("パスワード不一致：", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		c.Abort()
	} else {
		log.Println("パスワード一致。ログイン処理を続行します")
	}

	// ユーザ名が存在し、かつパスワードチェックもOKの場合、セッションIDとセッションユーザ名を発行
	session := sessions.Default(c)
	var user entity.User
	fmt.Println("session-id is ", session.Get("session-id"))

	//user.SessionId = session.Get("session-id" //こうするべき？でもsession-id入ってない
	user.Name = cookieUserName
	user.SessionId = cookieSessionId
	user = cont.DbConn.UpdateSessionID(&user)

	session.Set("user_name", dbUser.Name)
	session.Save()
	//}
	c.JSON(http.StatusOK, gin.H{"user name is*": session.Get("user_name")})
}

func (cont *Controller) BeforeAfter() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("------------before logic")
		session := sessions.Default(c)
		fmt.Println(session)
		cookieUserName, _ := c.Cookie("user_name")
		cookieSessionId, _ := c.Cookie("session-id")
		fmt.Println("cookieSessionId is", cookieSessionId)
		var dbUser entity.User
		_, dbUser = cont.DbConn.UserCount(cookieUserName)

		if dbUser.SessionId == cookieSessionId &&
			dbUser.Name == cookieUserName {
			log.Println("DBのセッションID&ユーザ名と、cookieの情報が一致しました。すでにログイン済です。")
			c.Set("session-id", cookieSessionId) //これいらない？
			c.Set("user-name", cookieUserName)
			// ユーザーの最終ログイン日時を更新する
			dbUser = entity.User{Name: cookieUserName, LastLoginAt: time.Now()}
			cont.DbConn.UpdateLastLoginAt(&dbUser)
			c.Next() // 以降、関数実行後の後処理
			log.Println("------------after logic2")
			c.Abort()
		} else {
			log.Println("DBのセッションID&ユーザ名と、cookieの情報が一致しません。ログインしてください。")
			// TODO: ログイン画面へ遷移
			c.Next() // 以降、関数実行後の後処理
			log.Println("------------after logic1")
			c.Abort()
		}
		log.Println("ログインチェック終わり")
	}
}

func (cont *Controller) PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	//セッションからデータを破棄する
	session := sessions.Default(c)
	session.Clear()
	session.Save()

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
