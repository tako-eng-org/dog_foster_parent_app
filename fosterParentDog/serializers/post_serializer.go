package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostResponse struct {
	ID             uint   `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Publishing     int    `json:"publishing"`      //公開設定
	DogName        string `json:"dog_name"`        //犬の名前
	Breed          string `json:"breed"`           //犬種
	Gender         int    `json:"gender"`          //性別
	Spay           int    `json:"spay"`            //去勢/避妊手術
	Old            string `json:"old"`             //年齢
	SinglePerson   int    `json:"single_person"`   //単身者への譲渡
	SeniorPerson   int    `json:"senior_person"`   //高齢者への譲渡
	TransferStatus int    `json:"transfer_status"` //譲渡ステータス
	Introduction   string `json:"introduction"`    //犬の自己紹介
	AppealPoint    string `json:"appeal_point"`    //性格アピールポイント
	OtherMessage   string `json:"other_message"`   //健康状態や譲渡条件などの特記事項
	UserId         uint   `json:"user_id"`         //ユーザーID
}

type PostSerializer struct {
	C    *gin.Context
	Post entity.Post
}

type PostsSerializer struct {
	C     *gin.Context
	Posts []entity.Post
}

func (ps *PostSerializer) Response() PostResponse {
	response := PostResponse{
		ID:             ps.Post.ID,
		CreatedAt:      ps.Post.CreatedAt.Format("2006-01-02 15:04"),
		UpdatedAt:      ps.Post.UpdatedAt.Format("2006-01-02 15:04"),
		Publishing:     ps.Post.Publishing,
		DogName:        ps.Post.DogName,
		Breed:          ps.Post.Breed,
		Gender:         ps.Post.Gender,
		Spay:           ps.Post.SeniorPerson,
		Old:            ps.Post.Old,
		SinglePerson:   ps.Post.SinglePerson,
		SeniorPerson:   ps.Post.SeniorPerson,
		TransferStatus: ps.Post.TransferStatus,
		Introduction:   ps.Post.Introduction,
		AppealPoint:    ps.Post.AppealPoint,
		OtherMessage:   ps.Post.OtherMessage,
		UserId:         ps.Post.UserId,
		//TopImagePath:   ps.Post.TopImagePath,
	}
	return response
}

func (ps *PostsSerializer) Response() []PostResponse {
	var response []PostResponse
	for _, post := range ps.Posts {
		serializer := PostSerializer{ps.C, post}
		response = append(response, serializer.Response())
	}
	return response
}
