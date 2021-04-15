package controller

import (
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 投稿idをもとに、投稿画像パスリストを取得する
//*******************************************************************
func (cont *Controller) FetchPostImagePathList(c *gin.Context) {
	model, _ := cont.DbConn.FindPostImagePathList(c.Query("postId"))
	c.Set("my_post_image_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	serializer := serializers.PostImagesSerializer{C: c, PostImages: model}
	c.JSON(http.StatusOK, serializer.Response())
}
