package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C          *gin.Context
	EntityUser entity.User
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
func (us *UserSerializer) ResponseUser() UserResponse {
	ret := UserResponse{
		ID:             us.EntityUser.ID,
		CreatedAt:      us.EntityUser.CreatedAt.Format("2006-01-02 15:04"),
		UpdatedAt:      us.EntityUser.UpdatedAt.Format("2006-01-02 15:04"),
		Name:           us.EntityUser.Name,
		Email:          us.EntityUser.Email,
		TwitterAccount: us.EntityUser.TwitterAccount,
		Nickname:       us.EntityUser.Nickname,
		Profile:        us.EntityUser.Profile,
		WebUrl:         us.EntityUser.WebUrl,
	}
	return ret
}
