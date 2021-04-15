package controller

import (
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 投稿idをもとに、譲渡可能都道府県idを取得する
//*******************************************************************
func (cont *Controller) FetchPostPrefectureList(c *gin.Context) {
	model, _ := cont.DbConn.FindPostPrefectureList(c.Query("postId"))
	c.Set("my_post_prefecture_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	serializer := serializers.PostPrefecturesSerializer{C: c, PostPrefectures: model}
	c.JSON(http.StatusOK, serializer.Response())
}
