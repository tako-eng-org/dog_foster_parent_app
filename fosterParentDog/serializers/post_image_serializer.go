package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostImageSerializer struct {
	C               *gin.Context
	EntityPostImage []entity.PostImage
}

type PostImageResponse struct {
	PostID      uint   `json:"post_id"`
	PostImageID uint   `json:"post_image_id"`
	ImagePath   string `json:"image_path"`
}

func (ps *PostImageSerializer) ResponsePostImage() []PostImageResponse {
	myModel := ps.C.MustGet("my_post_image_model").([]entity.PostImage)
	var postRet []PostImageResponse
	for i := 0; i < len(myModel); i++ {
		postRet = append(postRet, PostImageResponse{
			PostID:      ps.EntityPostImage[i].PostId,
			PostImageID: ps.EntityPostImage[i].ID,
			ImagePath:   ps.EntityPostImage[i].ImagePath,
		})
	}
	return postRet
}
