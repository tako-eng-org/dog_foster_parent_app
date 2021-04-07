package db

import (
	"fmt"
	entity "fpdapp/models/entity"
	serializer "fpdapp/serializers"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgres用ドライバ
)

//*******************************************************************
// DB接続する
//*******************************************************************
func open() *gorm.DB {
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

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.PostImage{},
		&entity.PostPrefecture{},
		&entity.Prefecture{},
	)

	fmt.Println("db connected: ", &db)
	return db
}

//*******************************************************************
// 公開済みレコードの数を取得する
//*******************************************************************
// return(ex): -> 41
func CountPublishedPostNum() int64 {
	var countNum int64

	db := open()
	//SELECT count(id) FROM "post"  WHERE (publishing = '0')
	db.Table("post").
		Select("count(id)").
		Where("publishing = ?", "0").
		Count(&countNum)

	defer db.Close()
	return countNum
}

//*******************************************************************
// 公開済み投稿を1ページ表示分取得する
//*******************************************************************
func FindIndexRecords(page string) ([]entity.Post, error) {
	db := open() // return *gorm.DB
	model := []entity.Post{}
	pageNum, _ := strconv.Atoi(page) // 数値に変換する
	numberPerPage := 20              // 1ページあたりの表示件数

	err := db.Order("id desc").
		Limit(numberPerPage).
		Offset((pageNum - 1) * numberPerPage).
		Find(&model). // ポインタを渡してメモリアドレスに結果を格納する
		Error
	defer db.Close()
	return model, err
}

//*******************************************************************
// [第1引数]の投稿IDでレコードを取得する
//*******************************************************************
func FindPostOneRecord(postIdStr string) []serializer.PostResponse {
	var records []serializer.PostResponse
	db := open()
	// 数値に変換する
	//postId, _ := strconv.Atoi(postIdStr)

	db.Where("id = ?", postIdStr).First(&records)
	defer db.Close()

	return records
}

//*******************************************************************
// [第1引数]の投稿IDで、投稿画像デーブルから投稿画像パスを取得する
//*******************************************************************
type Results struct {
	Id        uint
	ImagePath string //名前はテーブル定義と統一すれば取得できる。
}

func FindPostImagePaths(postIdStr string) []Results {
	var ResultUrls []Results
	db := open()

	// TODO 全部取れてない気がする
	db.Table("post").
		Select("post_image.id, post_image.image_path").
		Joins("join post_image on post.id = post_image.post_id").
		Where("post.id <> ?", postIdStr).
		Scan(&ResultUrls)

	defer db.Close()
	return ResultUrls
}

//*******************************************************************
// レコードを登録する
//*******************************************************************
func InsertRecord(registerRecord *entity.Post) {
	db := open()
	db.Create(&registerRecord) // insert
	defer db.Close()
}

////debug-----------------------start
//func Get20Records(page string) ([]entity.Post, error) {
//	db := open() // return *gorm.DB
//	model := []entity.Post{}
//	pageNum, _ := strconv.Atoi(page) // 数値に変換する
//	numberPerPage := 20              // 1ページあたりの表示件数
//
//	err := db.Order("id desc").
//		Limit(numberPerPage).
//		Offset((pageNum - 1) * numberPerPage).
//		Find(&model). // ポインタを渡してメモリアドレスに結果を格納する
//		Error
//	//fmt.Printf("%v", model)
//	defer db.Close()
//	return model, err
//}
////------------------
