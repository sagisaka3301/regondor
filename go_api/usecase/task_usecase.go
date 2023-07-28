package usecase

import (
	"go_api/model"
	"go_api/repository"
	"go_api/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

// taskUsecaseの構造体
// trというフィールド名でRepositoryパッケージ内のITaskRepositoryインターフェースの値を格納できるようにしておく。
// tvというフィールド名でTaskValidatorというフィールドを追加しておく。
type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

// NewTaskUsecaseのコンストラクター
// 外側でインスタンス化されているtaskValidatorを注入できるように引数にItaskValidatorを追加。
func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	// taskRepository, taskValidatorの機能をtaskUsecaseの中で使用できるようにしておく、
	return &taskUsecase{tr, tv} // アドレスを取得し返す。
}

// 返り値の一つ目の型として、taskResponse構造体の配列の型を指定しておく。
func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	// 取得するタスクの一覧を格納するためのTask構造体のスライスを定義。
	tasks := []model.Task{}
	// RepositoryにあるGetAllTasks()を呼び出す。
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	// 取得に成功した場合は、クライアントへのレスポンス用のTaskResponse構造体を0値で作成する。
	resTasks := []model.TaskResponse{}
	// for文でタスクを1つ1つ取り出してタスクレスポンス構造体を新しく作っていく。
	// 作成した新しい構造体をresTasksのスライスにappendで追加していく。
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	// resTasks外に返すようにしている。
	return resTasks, nil
}

// idからタスクを取り出す。
func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

// createTasks
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// リポジトリのCreateTaskを呼び出す前にtaskValidationを実行する。
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	// エラーが発生した場合、TaskResponse構造体の0値の実体とエラーをreturnで返す。
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	// 成功した場合は、引数で渡したアドレスが指し示す先の値が新規作成したタスクの値で書き変わる。
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	// リポジトリのCreateTaskを呼び出す前にtaskValidationを実行する。
	// バリデーションをかけたいtaskを引数に入れる。
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	// UpdateTaskでタスクオブジェクトのアドレスとユーザーID、タスクIDを渡している。
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
