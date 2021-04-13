package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostImageSerializer struct {
	C               *gin.Context
	EntityPostImage entity.PostImage
}

type PostImageResponse struct {
	PostID      uint   `json:"post_id"`
	PostImageID uint   `json:"post_image_id"`
	ImagePath   string `json:"image_path"`
}

func (ps *PostImageSerializer) Response() PostImageResponse {
	//myModel := ps.C.MustGet("my_post_image_model").(entity.PostImage)
	ret := PostImageResponse{
		PostID:      ps.EntityPostImage.PostId,
		PostImageID: ps.EntityPostImage.ID,
		ImagePath:   ps.EntityPostImage.ImagePath,
	}
	return ret
}
