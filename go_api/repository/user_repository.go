package repository

import (
	"go_api/model"

	"gorm.io/gorm"
)

// リポジトリのインターフェース
type IUserRepository interface {
	// 第一引数でユーザーオブジェクトのポインタ、第二引数で検索したいメールをstringで受け取れるようにしている。
	// 返り値はerrorインターフェース型
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// 実際のリポジトリの構造体
type userRepository struct {
	db *gorm.DB
}

// リポジトリにデータベースのインスタンスをDependency インジェクションするために、リポジトリの方にもコンストラクターを作成(依存性の注入)
// 外側でインスタンス化されたdbを引数で受け取る。
// 返り値の型はIUserRepositoryのインターフェース型に指定。
// 引数で受け取ったdbのインスタンスを要素にして、ユーザーリポジトリの構造体の実体を作成し、ポインターをリターンで返す。
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// 引数の型と返り値の型はインターフェースで定義したものと全く同じ必要がある。
// 引数でユーザーからの入力値を受け取る。(ユーザーオブジェクトで)
// オブジェクトはメタモン状態(空き箱)。メソッドによって内容が変わる。
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// めっちゃ重要：引数で受け取ったユーザーオブジェクトのアドレスが示す内容を検索したユーザーオブジェクトの内容に書き換える。
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// 引数で受け取ったユーザーオブジェクトのポインタを渡す。
	// めっちゃ重要：ユーザー作成に成功した場合、このポインタが指し示す値が新しく作成されたユーザーの情報に書き変わる。
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
