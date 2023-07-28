package repository

import (
	"fmt"
	"go_api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	// ログインしているユーザーが作成したタスクの一覧を取得するメソッド。
	// タスクの一覧を配列に格納するためにmodelタスクのスライスのポインタを第一引数で渡す。
	// 第2引数はログインしているユーザーのidを渡す。
	GetAllTasks(task *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// まずはtaskRepositoryという構造体を定義する。
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepositoryというコンストラクターを作成。
// 外側(上)でインスタンス化されたdbを引数として受け取って、受け取ったDBを使ってTaskRepositoryの構造体の実体を作成する。
// そして、その実体のアドレスを取得してreturnで返す。
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// ログイン済みのユーザーが作成したすべての投稿一覧を取得
// ブレークポイントとはソフトウェアのデバッグ中にプログラムの実行を一時停止するための指定されたポイント
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// 特定のidのタスクを取得するメソッド
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	// taskの主キー(id)が引数で受け取ったtaskIDに一致するtaskを取得する。
	// そして、取得したタスクオブジェクトを引数で受け取っていたポインタアドレスが指し示す先のメモリー領域に書き込む
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// CreateTaskメソッド
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTaskメソッド
// Clauses(clause.Returning{})をつけると、更新したあとのタスクのオブジェクトをこのタスクのポインタが指し示す先に書き込んでくれる。
// titleの値を引数で受け取るTaskオブジェクトのTitleの値に更新する。
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	// 処理の返り値をresultという変数に代入し、reslt.Errorでエラーを取得する。
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	// updateの場合、更新しようとしたオブジェクトが存在しないときはエラーにならない仕様になっている。
	// RowsAffectedで実際に更新されたレコードの数を取得することができ、その数が1より小さい(=0)の時は、更新が行われなかったことを意味するので、エラーを返す。
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// DeleteTask
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
