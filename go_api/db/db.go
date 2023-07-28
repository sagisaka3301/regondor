// dbを起動するためのパッケージ
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gormを使った場合、データベースを起動したときに、gormパッケージで定義されているDBという構造体の実態のアドレスが返ってくる。
// よって、関数の返り値をgorm.DBのポインタ型(*gorm.DB)する。
func NewDB() *gorm.DB {
	// 環境変数を読み込むための処理
	// 見た通り、Load()をする前にenvで条件判定を行っているので、GO_ENVだけはファイル実行時にターミナル上で渡す。
	// GO_ENV=dev コマンド
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// DBに接続するためのurlを作成。　復習：Sprintでstringを生成できる。
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db

}

// DBを閉じる(切断する)関数
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
