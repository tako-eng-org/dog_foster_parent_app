package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostPrefectureSerializer struct {
	C              *gin.Context
	PostPrefecture entity.PostPrefecture
}

type PostPrefecturesSerializer struct {
	C               *gin.Context
	PostPrefectures []entity.PostPrefecture
}

type PostPrefectureResponse struct {
	PostId           uint `json:"post_id"`
	PostPrefectureId int  `json:"post_prefecture_id"`
}

func (ps *PostPrefectureSerializer) Response() PostPrefectureResponse {
	response := PostPrefectureResponse{
		PostId:           ps.PostPrefecture.PostId,
		PostPrefectureId: ps.PostPrefecture.PostPrefectureId,
	}
	return response
}

func (ps *PostPrefecturesSerializer) Response() []PostPrefectureResponse {
	var response []PostPrefectureResponse
	for _, postPrefecture := range ps.PostPrefectures {
		serializer := PostPrefectureSerializer{ps.C, postPrefecture}
		response = append(response, serializer.Response())
	}
	return response
}
