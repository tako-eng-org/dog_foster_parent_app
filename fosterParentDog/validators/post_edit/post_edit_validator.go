package post_edit

import (
	"fmt"
	"fpdapp/models/entity"

	"github.com/go-playground/validator/v10"
)

// 受け取るリクエストの構造体
type Request struct {
	Publishing     int    `json:"publishing" validate:"gte=0,lt=2"`
	DogName        string `json:"dog_name"`
	Breed          string `json:"breed"`
	Gender         int    `json:"gender"`
	Spay           int    `json:"spay"`
	Old            string `json:"old"`
	SinglePerson   int    `json:"single_person"`
	SeniorPerson   int    `json:"senior_person"`
	TransferStatus int    `json:"transfer_status"`
	Introduction   string `json:"introduction"`
	AppealPoint    string `json:"appeal_point"`
	OtherMessage   string `json:"other_message"`
	UserId         uint   `json:"user_id"`

	PostImages      []PostImageRequest      `json:"post_images"`
	PostPrefectures []PostPrefectureRequest `json:"post_prefectures"`
}
type PostImageRequest struct {
	// TODO: 要実装/Updateの際は必要になるのでコメントにしている
	//PostId      uint   `json:"post_id"`
	Position int            `json:"position"`
	ImageId  uint           `json:"image_id"`
	Images   []ImageRequest `json:"images"`
}
type ImageRequest struct {
	ObjectKey string
	UserId    uint
}

type PostPrefectureRequest struct {
	//PostId           uint `json:"post_id"`
	PostPrefectureId int `json:"post_prefecture_id"`
}

type Validator struct {
	Post Request
}

// TODO: 戻り値など、この構成でOKか要確認
func (r *Validator) Request() *entity.Post {
	var postImageRequestList []entity.PostImage
	for _, p := range r.Post.PostImages {
		postImageRequestList = append(postImageRequestList, entity.PostImage{
			Position: p.Position,
			// TODO: requestへ画像投稿にまつわる情報を追加する
		})
	}

	var postPrefectureRequestList []entity.PostPrefecture
	for _, p := range r.Post.PostPrefectures {
		postPrefectureRequestList = append(postPrefectureRequestList, entity.PostPrefecture{
			PostPrefectureId: p.PostPrefectureId,
		})
	}

	request := &entity.Post{
		Publishing:      r.Post.Publishing,
		DogName:         r.Post.DogName,
		Breed:           r.Post.Breed,
		Gender:          r.Post.Gender,
		Spay:            r.Post.SeniorPerson,
		Old:             r.Post.Old,
		SinglePerson:    r.Post.SinglePerson,
		SeniorPerson:    r.Post.SeniorPerson,
		TransferStatus:  r.Post.TransferStatus,
		Introduction:    r.Post.Introduction,
		AppealPoint:     r.Post.AppealPoint,
		OtherMessage:    r.Post.OtherMessage,
		UserId:          r.Post.UserId,
		PostImages:      postImageRequestList,
		PostPrefectures: postPrefectureRequestList,
	}

	// バリデーションチェック
	// FIXME: バリデーションチェックが効いていない
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}
	return request
}
