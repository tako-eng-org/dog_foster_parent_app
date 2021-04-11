package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

//PostImage-------------------------------
type UserSerializer struct {
	C          *gin.Context
	EntityUser []entity.User
}

type UserResponse struct {
	ID             uint   `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Name           string `json:"post_id"`         //ユーザー名
	Email          string `json:"e-mail"`          //メールアドレス
	TwitterAccount string `json:"twitter_account"` //twitterアカウント
	Nickname       string `json:"nickname"`        //ニックネーム
	Profile        string `json:"profile"`         //自己紹介文
	WebUrl         string `json:"web_url"`         //websiteやSNSなどのURL
}

// *PostSerializer型の構造体 selfのレシーバを使用して、メソッドを定義する
func (us *UserSerializer) ResponseUser() []UserResponse {
	myModel := us.C.MustGet("my_post_user_model").([]entity.User)
	var postRet []UserResponse
	for i := 0; i < len(myModel); i++ {
		postRet = append(postRet, UserResponse{
			ID:             us.EntityUser[i].ID,
			CreatedAt:      us.EntityUser[i].CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:      us.EntityUser[i].UpdatedAt.Format("2006-01-02 15:04"),
			Name:           us.EntityUser[i].Name,
			Email:          us.EntityUser[i].Email,
			TwitterAccount: us.EntityUser[i].TwitterAccount,
			Nickname:       us.EntityUser[i].Nickname,
			Profile:        us.EntityUser[i].Profile,
			WebUrl:         us.EntityUser[i].WebUrl,
		})
	}
	return postRet
}
