package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

//単数
type PostPrefectureSerializer struct {
	C              *gin.Context
	PostPrefecture entity.PostPrefecture
}

//複数
type PostPrefecturesSerializer struct {
	C               *gin.Context
	PostPrefectures []entity.PostPrefecture
}

type PostPrefectureResponse struct {
	PostId           uint `json:"post_id"`
	PostPrefectureId int  `json:"post_prefecture_id"`
}

//単数
func (ps *PostPrefectureSerializer) Response() PostPrefectureResponse {
	response := PostPrefectureResponse{
		PostId:           ps.PostPrefecture.PostId,
		PostPrefectureId: ps.PostPrefecture.PostPrefectureId,
	}
	return response
}

//複数
func (ps *PostPrefecturesSerializer) Response() []PostPrefectureResponse {
	var response []PostPrefectureResponse
	for _, postPrefecture := range ps.PostPrefectures {
		serializer := PostPrefectureSerializer{ps.C, postPrefecture}
		response = append(response, serializer.Response())
	}
	return response
}
