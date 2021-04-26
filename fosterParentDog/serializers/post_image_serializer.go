package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostImageSerializer struct {
	C         *gin.Context
	PostImage entity.PostImage
}

type PostImagesSerializer struct {
	C          *gin.Context
	PostImages []entity.PostImage
}

type PostImageResponse struct {
	PostID      uint   `json:"post_id"`
	PostImageID uint   `json:"post_image_id"`
	ImagePath   string `json:"image_path"`
}

func (ps *PostImageSerializer) Response() PostImageResponse {
	response := PostImageResponse{
		PostID:      ps.PostImage.PostId,
		PostImageID: ps.PostImage.ID,
		ImagePath:   ps.PostImage.ImagePath,
	}
	return response
}

func (ps *PostImagesSerializer) Response() []PostImageResponse {
	var response []PostImageResponse
	for _, image := range ps.PostImages {
		serializer := PostImageSerializer{ps.C, image}
		response = append(response, serializer.Response())
	}
	return response
}
