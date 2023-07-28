package main

import (
	"fmt"
	"go_api/db"
	"go_api/model"
)

// マイグレーションファイルはmainパッケージに所属させる。
// マイグレーションを実行したい場合は、gorm.DBをポインタレシーバとして、オートマイグレートというメソッドが実装されているのでこれを呼び出す。
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// 引数に、データベースに反映させたいモデル構造を渡す。
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
