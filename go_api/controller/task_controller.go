package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

// 構造体を定義
// tuというフィールド名でusecaseパッケージ内のtaskusecaseインターフェースの値を格納できるようにしておく。
type taskController struct {
	tu usecase.ITaskUsecase
}

// NewTaskControllerのコンストラクターは引数で外側似てインスタンス化されているタスクのusecaseを受け取る。
// 受け取った値を使ってタスクコントローラーの構造体のインスタンスを作成し、そのポインターをリターンで返すようにしておく。
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	// まず、ユーザーから送られてくるjwtトークンに組み込まれているユーザーIDの値を取り出す。
	// routerに実装するJWTのミドルウェア側で送られてきたJWTトークンをデコードしてくれる。
	// そしてデコードした内容をechoのコンテキストの中にuserというフィールド名をつけて自動的に格納してくれる。
	// コントローラー側では、そのuserというキーワードを使って、コンテキストからjwtをデコードした値を読み込む。
	// その中には、デコードされたclaimsが格納されているので、user.Claimsで取り出し、Claimsの中にあるuser_idを取得し、変数に代入。
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// コンテキストから取得した値はany型になっているので、いったんfloat64に型アサーションしてからuintに型変換する。
	// そして、taskUsecaseのGetalltasksメソッドにユーザーidを引数として渡す。
	// エラーが発生した場合、context.Jsonでクライアントにinternalservererrorとエラーメッセージを返す。
	taskRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	// userキーワードを使ってjwtをデコードした内容を取得。
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// その中からuseridの値を取得し、変数に格納。
	// このuser_idとはjwtの一部としてエンコードされており、リクエストを送信したuserのidを表している。
	// よって、これはtaskではなく、ログインしているuser自体のidを示す。
	userId := claims["user_id"]
	// リクエストパラメーターからtaskIdを取得する。
	// string型になっているので、
	id := c.Param("taskId")
	// Atoi(エートゥアイ)を使用してstringからint型に変更。
	taskId, _ := strconv.Atoi(id)
	// usecaseのgetTaskByIDメソッドを呼び出す。
	// 第一引数にuser_id,第二引数にtaskIdを渡す。
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	// エラーが発生した場合、internalServerError、成功した場合は、StatusOKで取得したタスクをJSONでクライアントに返す。
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	// コンテキストからユーザーIDを取得。
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// 0値でタスク構造体を作っておき、context.Bindを使うことでリクエストbodyに含まれる内容をタスク構造体に代入する。
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// タスクオブジェクトのユーザーidのフィールドにコンテキストから取得したユーザーidの値を格納する。
	task.UserId = uint(userId.(float64))
	// そのタスクオブジェクトをtaskusecaseのCreateTaskに引数として渡す。
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	// コンテキストからユーザーidを取得する。
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	// リクエストパラメーターからtaskIdを取得する。
	id := c.Param("taskId")
	// stringからint型に変化する。
	taskId, _ := strconv.Atoi(id)

	// コンテキストBindを使って、リクエストオブジェクトの値をタスクオブジェクトにバインドする。
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// タスクusecaseのupdateTaskを呼び出す。第一引数：userId、第二引数：taskId
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 成功した場合、更新後のタスクの値をステータスOKでクライアントにjsonで返す。
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// NoContentに関する解説はなかった。
	return c.NoContent(http.StatusNoContent)
}
