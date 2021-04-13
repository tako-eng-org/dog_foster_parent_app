package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostSerializer struct {
	C          *gin.Context
	EntityPost entity.Post
}

type PostResponse struct {
	ID               uint   `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	Publishing       int    `json:"publishing"`         //公開設定
	DogName          string `json:"dog_name"`           //犬の名前
	Breed            string `json:"breed"`              //犬種
	Gender           int    `json:"gender"`             //性別
	Spay             int    `json:"spay"`               //去勢/避妊手術
	Old              string `json:"old"`                //年齢
	SinglePerson     int    `json:"single_person"`      //単身者への譲渡
	SeniorPerson     int    `json:"senior_person"`      //高齢者への譲渡
	TransferStatus   int    `json:"transter_status"`    //譲渡ステータス
	Introduction     string `json:"introduction"`       //犬の自己紹介
	AppealPoint      string `json:"appeal_point"`       //性格アピールポイント
	PostPrefectureId int    `json:"post_prefecture_id"` //譲渡可能都道府県ID
	OtherMessage     string `json:"other_message"`      //健康状態や譲渡条件などの特記事項
	UserId           uint64 `json:"user_id"`            //ユーザーID
	TopImagePath     string `json:"top_image_path"`     //top投稿画像パス
	PostImageId      uint64 `json:"post_image_id"`      //投稿画像ID
}

func (ps *PostSerializer) Response() PostResponse {
	//myPostModel := ps.C.MustGet("my_post_model").([]entity.Post)
	//var retSlice PostResponse
	//for i := 0; i < len(myPostModel); i++ {
	ret := PostResponse{
		ID:               ps.EntityPost.ID,
		CreatedAt:        ps.EntityPost.CreatedAt.Format("2006-01-02 15:04"),
		UpdatedAt:        ps.EntityPost.UpdatedAt.Format("2006-01-02 15:04"),
		Publishing:       ps.EntityPost.Publishing,
		DogName:          ps.EntityPost.DogName,
		Breed:            ps.EntityPost.Breed,
		Gender:           ps.EntityPost.Gender,
		Spay:             ps.EntityPost.SeniorPerson,
		Old:              ps.EntityPost.Old,
		SinglePerson:     ps.EntityPost.SinglePerson,
		SeniorPerson:     ps.EntityPost.SeniorPerson,
		TransferStatus:   ps.EntityPost.TransferStatus,
		Introduction:     ps.EntityPost.Introduction,
		AppealPoint:      ps.EntityPost.AppealPoint,
		PostPrefectureId: ps.EntityPost.PostPrefectureId,
		OtherMessage:     ps.EntityPost.OtherMessage,
		UserId:           ps.EntityPost.UserId,
		TopImagePath:     ps.EntityPost.TopImagePath,
		PostImageId:      ps.EntityPost.PostImageId,
	}
	//}
	return ret
}
