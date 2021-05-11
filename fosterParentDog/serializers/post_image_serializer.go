package serializers

import (
	"fpdapp/models/entity"

	"github.com/gin-gonic/gin"
)

type PostImageResponse struct {
	PostID      uint   `json:"post_id"`
	PostImageID uint   `json:"post_image_id"`
	Position    int    `json:"position"`
	ImagePath   string `json:"image_path"`
}

type PostImageSerializer struct {
	C         *gin.Context
	PostImage entity.PostImage
}

type PostImageListSerializer struct {
	C          *gin.Context
	PostImages []entity.PostImage
}

func (ps *PostImageSerializer) Response() PostImageResponse {
	response := PostImageResponse{
		PostID:      ps.PostImage.PostId,
		PostImageID: ps.PostImage.ID,
		Position:    ps.PostImage.Position,
		//ImagePath:   ps.PostImage.ImagePath,
	}
	return response
}

func (ps *PostImageListSerializer) Response() []PostImageResponse {
	var response []PostImageResponse
	for _, image := range ps.PostImages {
		serializer := PostImageSerializer{ps.C, image}
		response = append(response, serializer.Response())
	}
	return response
}
