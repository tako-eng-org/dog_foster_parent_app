package db

import (
	"fpdapp/models/entity"
	"log"
)

//*******************************************************************
// [第1引数]の投稿IDで、投稿画像テーブルから投稿画像パスを取得する
//*******************************************************************
func (db *Database) FindPostImagePathList(postId string) []entity.PostImage {
	var model []entity.PostImage

	//SELECT post.id as post_id,
	//post_image.id as post_image_id
	//post_image.image_path
	//FROM "post"
	//left join post_image
	//on post.id = post_image.post_id
	//WHERE (post.id = '44')
	err := db.connection.Table("posts").
		Select("posts.id as post_id,"+
			" post_images.id as post_image_id,"+
			" post_images.image_path").
		Joins("left join post_images"+
			" on posts.id = post_images.post_id ").
		Where("posts.id = ?", postId).
		Scan(&model).
		Error
	if err != nil {
		log.Fatal(err)
	}

	return model
}
