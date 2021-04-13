package db

import (
	"fmt"
	"fpdapp/models/entity"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgres用ドライバ
)

type (
	Database struct {
		db *gorm.DB
	}
)

//*******************************************************************
// DB接続する
//*******************************************************************
func Open() *Database {
	// .envファイルから環境変数を読み出す
	fileEnv := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))

	// .envの読み込みに失敗した場合
	if fileEnv != nil {
		panic(fileEnv)
	}

	// DB接続
	db, err := gorm.Open(os.Getenv("RDBMS_NAME"),
		"host= "+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USER")+
			" password="+os.Getenv("DB_PASS")+
			" dbname="+os.Getenv("DB_NAME")+
			" sslmode="+os.Getenv("DB_SSLMODE"))

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を複数形で扱う（デフォルトは複数形）
	db.SingularTable(false)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.PostImage{},
		&entity.PostPrefecture{},
	)

	fmt.Println("db connected: ", &db)
	return &Database{db: db}
}

//*******************************************************************
// DBをクローズする
//*******************************************************************
func (db *Database) Close() {
	_ = db.db.Close()
}

//*******************************************************************
// 公開済みレコードの数を取得する
//*******************************************************************
// return(ex): -> 41
func (db *Database) CountPublishedPostNum() int64 {
	var countNum int64

	//SELECT count(id) FROM "post"  WHERE (publishing = '0')
	db.db.Table("posts").
		Select("count(id)").
		Where("publishing = ?", "0").
		Count(&countNum)

	return countNum
}

//*******************************************************************
// 公開済み投稿を1ページ表示分取得する
//*******************************************************************
func (db *Database) FindIndexRecords(page string) ([]entity.Post, error) {

	var model []entity.Post
	pageNum, _ := strconv.Atoi(page) // 数値に変換する
	numberPerPage := 20              // 1ページあたりの表示件数

	err := db.db.Order("id desc").
		Limit(numberPerPage).
		Offset((pageNum - 1) * numberPerPage).
		Find(&model). // ポインタを渡してメモリアドレスに結果を格納する
		Error
	return model, err
}

//*******************************************************************
// [第1引数]の投稿IDでレコードを取得する
//*******************************************************************
func (db *Database) FindPostOneRecord(postId string) ([]entity.Post, error) {
	var model []entity.Post

	err := db.db.Where("id = ?", postId).First(&model).Error

	return model, err
}

//*******************************************************************
// [第1引数]の投稿IDで、投稿画像テーブルから投稿画像パスを取得する
//*******************************************************************
func (db *Database) FindPostImagePaths(postIdStr string) ([]entity.PostImage, error) {
	var model []entity.PostImage

	//SELECT post.id as post_id,
	//post_image.id as post_image_id
	//post_image.image_path
	//FROM "post"
	//left join post_image
	//on post.id = post_image.post_id
	//WHERE (post.id = '44')
	err := db.db.Table("posts").
		Select("posts.id as post_id,"+
			" post_images.id as post_image_id,"+
			" post_images.image_path").
		Joins("left join post_images"+
			" on posts.id = post_images.post_id ").
		Where("posts.id = ?", postIdStr).
		Scan(&model).
		Error

	if err != nil {
		fmt.Println(err)
	}
	return model, err
}

//*******************************************************************
// [第1引数]の投稿IDで、譲渡可能都道府県を取得する
//*******************************************************************
func (db *Database) FindPostFetchPostTransferablePrefecture(postIdStr string) ([]entity.TransferablePrefecture, error) {
	var model []entity.PostPrefecture

	//select post_prefecture.id, post_prefecture.post_prefecture_id from post_prefecture where post_id = 44;
	err := db.db.Table("post_prefectures").
		Select("post_prefectures.id,"+
			" post_prefectures.post_id,"+
			" post_prefectures.post_prefecture_id").
		Where("post_prefectures.post_id = ?", postIdStr).
		Scan(&model).
		Error

	if err != nil {
		fmt.Println(err)
	}
	return model, err
}

//*******************************************************************
// [第1引数]の投稿IDで、ユーザー情報を取得する
//*******************************************************************
func (db *Database) FindPostUser(postIdStr string) (entity.User, error) {
	var model entity.User

	/*
		select B.*
		from post A
		inner join public.user B
		on A.user_id = B.id
		where A.id = 44;
	*/
	err := db.db.Table("posts A").
		Select("B.*").
		Joins("inner join public.users B on A.user_id = B.id").
		Where("A.id = ?", postIdStr).
		Scan(&model).
		Error

	if err != nil {
		fmt.Println(err)
	}
	return model, err
}

//*******************************************************************
// レコードを登録する
//*******************************************************************
func (db *Database) InsertRecord(registerRecord *entity.Post) {
	db.db.Create(&registerRecord) // insert
}
