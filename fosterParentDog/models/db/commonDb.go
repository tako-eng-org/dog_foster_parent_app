package db

import (
	"fmt"
	entity "fpdapp/models/entity"
	"os"

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
