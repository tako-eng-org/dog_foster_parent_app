package controller

import (
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 投稿idをもとに、ユーザー情報を取得する
//*******************************************************************
func (pc *Controller) FetchPostUser(c *gin.Context) {
	postId := c.Query("postId")
	model, _ := pc.Database.FindPostUser(postId) //ORMを叩いてデータとerrを取得する

	c.Set("my_post_user_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.UserSerializer{C: c, EntityUser: model}
	c.JSON(http.StatusOK, posts.ResponseUser())
}
