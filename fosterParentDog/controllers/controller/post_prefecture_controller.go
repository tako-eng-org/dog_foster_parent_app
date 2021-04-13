package controller

import (
	"fpdapp/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*******************************************************************
// 投稿idをもとに、譲渡可能都道府県idを取得する
//*******************************************************************
func (cont *Controller) FetchPostPrefecture(c *gin.Context) {
	postId := c.Query("postId")
	model, _ := cont.DbConn.FindPostPrefecture(postId) //ORMを叩いてデータとerrを取得する

	// FIXME: シリアライザーを通さず処理しており、期待動作するが要修正。
	var ret []serializers.PostPrefectureResponse
	for i := 0; i < len(model); i++ {
		ret = append(ret, serializers.PostPrefectureResponse{
			PostId:           model[i].PostId,
			PostPrefectureId: model[i].PostPrefectureId,
		})
	}
	c.Set("my_post_prefecture_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
	//posts := serializers.PostPrefectureSerializer{C: c, EntityPostPrefecture: model}
	c.JSON(http.StatusOK, ret)
}

/*
   func (cont *Controller) FetchPostPrefecture(c *gin.Context) {
   	postId := c.Query("postId")
   	model, _ := cont.DbConn.FindPostPrefecture(postId) //ORMを叩いてデータとerrを取得する

   	c.Set("my_post_prefecture_model", model) //回避 (*Context).MustGet: panic("Key \"" + key + "\" does not exist")
   	posts := serializers.PostPrefectureSerializer{C: c, EntityPostPrefecture: model}
   	c.JSON(http.StatusOK, posts.ResponsePostPrefecture())
   }
*/
