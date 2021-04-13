package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostPrefectureSerializer struct {
	C                    *gin.Context
	EntityPostPrefecture entity.PostPrefecture
}

type PostPrefectureResponse struct {
	PostId           uint `json:"post_id"`
	PostPrefectureId int  `json:"post_prefecture_id"`
}

func (ps *PostPrefectureSerializer) ResponsePostPrefecture() PostPrefectureResponse {
	//myModel := ps.C.MustGet("my_post_prefecture_model").(entity.PostPrefecture)
	ret := PostPrefectureResponse{
		PostId:           ps.EntityPostPrefecture.PostId,
		PostPrefectureId: ps.EntityPostPrefecture.PostPrefectureId,
	}
	return ret
}

/*
func (ps *PostPrefectureSerializer) ResponsePostPrefecture() []PostPrefectureResponse {
	myModel := ps.C.MustGet("my_post_prefecture_model").([]entity.PostPrefecture)
	var postRet []PostPrefectureResponse
	for i := 0; i < len(myModel); i++ {
		postRet = append(postRet, PostPrefectureResponse{
			PostId:           ps.EntityPostPrefecture[i].PostId,
			PostPrefectureId: ps.EntityPostPrefecture[i].PostPrefectureId,
		})
	}
	return postRet
}
*/
