package serializers

import (
	entity "fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

// 配列
type EntityPost []entity.Post

type PostSerializer struct {
	C *gin.Context
	EntityPost
}

type PostResponse struct {
	ID                     uint   `json:"id"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	Publishing             int    `json:"publishing"`              //公開設定
	DogName                string `json:"dog_name"`                //犬の名前
	Breed                  string `json:"breed"`                   //犬種
	Gender                 int    `json:"gender"`                  //性別
	Spay                   int    `json:"spay"`                    //去勢/避妊手術
	Old                    string `json:"old"`                     //年齢
	SinglePerson           int    `json:"single_person"`           //単身者への譲渡
	SeniorPerson           int    `json:"senior_person"`           //高齢者への譲渡
	TransferStatus         int    `json:"transter_status"`         //譲渡ステータス
	Introduction           string `json:"introduction"`            //犬の自己紹介
	AppealPoint            string `json:"appeal_point"`            //性格アピールポイント
	TransferablePrefecture int    `json:"transferable_prefecture"` //譲渡可能都道府県ID
	OtherMessage           string `json:"other_message"`           //健康状態や譲渡条件などの特記事項
	UserId                 uint64 `json:"user_id"`                 //ユーザーID
	TopImagePath           string `json:"top_image_path"`          //top投稿画像パス
	PostImageId            uint64 `json:"post_image_id"`           //投稿画像ID
}

// *PostSerializer型の構造体 selfのレシーバを使用して、メソッドを定義する
func (self *PostSerializer) Response() []PostResponse {
	myPostModel := self.C.MustGet("my_post_model").([]entity.Post)
	postRet := []PostResponse{}
	for i := 0; i < len(myPostModel); i++ {
		postRet = append(postRet, PostResponse{
			ID:                     myPostModel[i].ID,
			CreatedAt:              myPostModel[i].CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:              myPostModel[i].UpdatedAt.Format("2006-01-02 15:04"),
			Publishing:             myPostModel[i].Publishing,
			DogName:                myPostModel[i].DogName,
			Breed:                  myPostModel[i].Breed,
			Gender:                 myPostModel[i].Gender,
			Spay:                   myPostModel[i].SeniorPerson,
			Old:                    myPostModel[i].Old,
			SinglePerson:           myPostModel[i].SinglePerson,
			SeniorPerson:           myPostModel[i].SeniorPerson,
			TransferStatus:         myPostModel[i].TransferStatus,
			Introduction:           myPostModel[i].Introduction,
			AppealPoint:            myPostModel[i].AppealPoint,
			TransferablePrefecture: myPostModel[i].TransferablePrefecture,
			OtherMessage:           myPostModel[i].OtherMessage,
			UserId:                 myPostModel[i].UserId,
			TopImagePath:           myPostModel[i].TopImagePath,
			PostImageId:            myPostModel[i].PostImageId,
		})
	}
	return postRet
}
