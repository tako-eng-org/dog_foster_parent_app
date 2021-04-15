package controller

import (
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 投稿idをもとに、ユーザー情報を取得する
//*******************************************************************
func (cont *Controller) FetchPostUser(c *gin.Context) {
	model, _ := cont.DbConn.FindPostUser(c.Query("postId"))
	c.Set("my_post_user_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	serializer := serializers.UserSerializer{C: c, User: model}
	c.JSON(http.StatusOK, serializer.ResponseUser())
}
