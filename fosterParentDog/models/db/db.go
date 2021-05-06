package db

import (
	"fmt"
	"fpdapp/models/entity"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgres用ドライバ
)

type Database struct {
	connection *gorm.DB
}

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
	conn, err := gorm.Open(os.Getenv("RDBMS_NAME"),
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
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	conn.LogMode(true)

	// 登録するテーブル名を複数形で扱う
	conn.SingularTable(false)

	// マイグレーション（テーブルが無い時は自動生成）
	conn.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.PostImage{},
		&entity.PostPrefecture{},
	)

	fmt.Println("db connected: ", &conn)
	return &Database{connection: conn}
}

//*******************************************************************
// DB接続をクローズする
//*******************************************************************
func (db *Database) Close() {
	_ = db.connection.Close()
}
