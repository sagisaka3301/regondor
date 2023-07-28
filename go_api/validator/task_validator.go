package validator

import (
	"go_api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// インターフェースを作成
type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

// 構造体を作成
type taskValidator struct{}

// コンストラクターを作成
// taskValidator構造体のインスタンスを作成するためのコンストラクターをNewTaskValidatorという名前で作成。
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// taskValidator型がITaskValidatorインターフェースを満たすために、TaskValidator型に対してtaskValidateメソッドを追加する。
// taskValidatorをポインターレシーバーとして受け取る形で定義。
// 引数で、バリデーションで評価したいTaskのオブジェクトを受け取る。
func (tv *taskValidator) TaskValidate(task model.Task) error {
	// validationStruct関数を実行。第一引数にtaskオブジェクトのアドレス。第二引数にタスクのタイトルに対するバリデーションを実装している。
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 12).Error("limited max 12 char"),
		),
	)
}
