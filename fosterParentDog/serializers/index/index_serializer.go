package index

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID              uint                     `json:"id"`
	CreatedAt       string                   `json:"created_at"`
	Publishing      int                      `json:"publishing"`
	DogName         string                   `json:"dog_name"`
	Breed           string                   `json:"breed"`
	Gender          int                      `json:"gender"`
	TransferStatus  int                      `json:"transfer_status"`
	Introduction    string                   `json:"introduction"`
	PostImages      []PostImageResponse      `json:"post_images"`
	PostPrefectures []PostPrefectureResponse `json:"post_prefectures"`
	User            UserResponse             `json:"user"`
}
type PostImageResponse struct {
	PostId      uint   `json:"post_id"`
	PostImageId uint   `json:"post_image_id"`
	Position    int    `json:"position"`
	ImagePath   string `json:"image_path"`
}
type PostPrefectureResponse struct {
	PostId           uint `json:"post_id"`
	PostPrefectureId int  `json:"post_prefecture_id"`
}
type UserResponse struct {
	Nickname string `json:"nickname"` //ニックネーム
	Profile  string `json:"profile"`  //自己紹介文
	WebUrl   string `json:"web_url"`  //websiteやSNSなどのURL
}

type Serializer struct {
	C    *gin.Context
	Post entity.Post
}

func (r *Serializer) Response() Response {
	var postImageResponseList []PostImageResponse
	for _, p := range r.Post.PostImages {
		postImageResponseList = append(postImageResponseList, PostImageResponse{
			PostId:      p.PostId,
			PostImageId: p.ID,
			Position:    p.Position,
			ImagePath:   p.ImagePath,
		})
	}

	var postPrefectureResponseList []PostPrefectureResponse
	for _, p := range r.Post.PostPrefectures {
		postPrefectureResponseList = append(postPrefectureResponseList, PostPrefectureResponse{
			PostId:           p.PostId,
			PostPrefectureId: p.PostPrefectureId,
		})
	}

	response := Response{
		ID:             r.Post.ID,
		CreatedAt:      r.Post.CreatedAt.Format("2006-01-02 15:04"),
		Publishing:     r.Post.Publishing,
		DogName:        r.Post.DogName,
		Breed:          r.Post.Breed,
		Gender:         r.Post.Gender,
		TransferStatus: r.Post.TransferStatus,
		Introduction:   r.Post.Introduction,

		PostImages:      postImageResponseList,
		PostPrefectures: postPrefectureResponseList,
		User: UserResponse{
			Nickname: r.Post.User.Nickname,
			Profile:  r.Post.User.Profile,
			WebUrl:   r.Post.User.WebUrl,
		},
	}
	return response
}

//type ListSerializer struct {
//	C     *gin.Context
//	Posts []entity.Post
//}
//
//func (ps *ListSerializer) Response() []Response {
//	var response []Response
//	for _, post := range ps.Posts {
//		serializer := Serializer{ps.C, post}
//		response = append(response, serializer.Response())
//	}
//	return response
//}
