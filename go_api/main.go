package main

import (
	"go_api/controller"
	"go_api/db"
	"go_api/repository"
	"go_api/router"
	"go_api/usecase"
	"go_api/validator"
)

func main() {
	// まず、DBをインスタンス化 dbパッケージで作成したNewDBメソッドを実行し、できたインスタンスを変数dbに格納。
	db := db.NewDB()
	// validatorのコンストラクターを実行し、構造体のインスタンスを作成する。
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	// リポジトリで作ったコンストラクターを起動する。 repositoryパッケージで作成したものを実行する。インスタンス化してあるdbを引数として注入。
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	mypageRepository := repository.NewMypageRepository(db)
	// ユースケースとタスクのコンストラクターも起動する。userRepositoryを引数にする。
	// validatorのインスタンスをユースケースのコンストラクターに渡す。
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	mypageUsecase := usecase.NewMypageUsecase(mypageRepository)
	// コントローラーのコンストラクターを起動する。userUsecase, taskUsecaseのインスタンスを引数として注入
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	mypageController := controller.NewMypageContorller(mypageUsecase)
	// routerの呼び出し。コントローラーを引数として注入。
	e := router.NewRouter(userController, taskController, mypageController)
	// echoインスタンスを使用し、サーバーを起動する。
	// e.Startで起動できる。ポートは8080。エラーが発生したとき、echoのLogger機能を使いログ情報を出力した後にプログラムを強制終了する。
	e.Logger.Fatal(e.Start(":8080"))

	// dependency injectionを追加。

}
