package controller

import (
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 投稿idをもとに、投稿画像パスを取得する
//*******************************************************************
func (pc *Controller) FetchPostImagePaths(c *gin.Context) {
	postId := c.Query("postId")
	postImageModel, _ := pc.Database.FindPostImagePaths(postId) //ORMを叩いてデータとerrを取得する

	c.Set("my_post_image_model", postImageModel) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	posts := serializers.PostImageSerializer{C: c, EntityPostImage: postImageModel}
	c.JSON(http.StatusOK, posts.ResponsePostImage())
}
