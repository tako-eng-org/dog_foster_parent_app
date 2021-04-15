package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C    *gin.Context
	User entity.User
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

func (us *UserSerializer) ResponseUser() UserResponse {
	response := UserResponse{
		ID:             us.User.ID,
		CreatedAt:      us.User.CreatedAt.Format("2006-01-02 15:04"),
		UpdatedAt:      us.User.UpdatedAt.Format("2006-01-02 15:04"),
		Name:           us.User.Name,
		Email:          us.User.Email,
		TwitterAccount: us.User.TwitterAccount,
		Nickname:       us.User.Nickname,
		Profile:        us.User.Profile,
		WebUrl:         us.User.WebUrl,
	}
	return response
}
