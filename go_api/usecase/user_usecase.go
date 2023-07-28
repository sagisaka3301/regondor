package usecase

import (
	"go_api/model"
	"go_api/repository"
	"go_api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	// ユーザーモデルをポインタではなく、値で受け取る。
	// 返り値の1つ目は、モデルで定義したUserResponse型にしている。
	// 返り値の2つ目は、エラーインターフェース型にしている。
	SignUp(user model.User) (model.UserResponse, error)
	// ログイン
	// 返り値の1つ目は、JWTトークンを返すためにstring型を割り当てている。2つ目はerrorインターフェース型にしている。
	Login(user model.User) (string, error)
}

// 構造体を定義
// フィールドとして、USERリポジトリを追加しておく
// usecaseのコードはリポジトリのインターフェースにだけ依存させるので、リポジトリパッケージで定義されているIUserRepositoryを使用。
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// usecaseにリポジトリをDEPENDENCYインジェクションするためのConstructorを書いておく
// 外部でインスタンス化されるリポジトリを引数で受け取るようにする。
// usecaseのコードはリポジトリのインターフェースだけに依存するので、引数の型はパッケージに定義されているIUserRepositoryのインターフェースの型を定義している。
// 引数で受け取れるリポジトリのインスタンスをフィールドとして、ユーザーユースケースの構造体の実体を作成する。
// 返り値の方はIUserUsecaseのインターフェース型
// そして、返り値で定義してるuserUsecaseがIUserUsecaseのインターフェースを満たす必要があるので、userUsecaseに対して、SignupとLoginを実装する必要がある。
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	// 作成した実体のポインタを&で取得してreturnで返す
	return &userUsecase{ur, uv}
}

// userusecaseをポインタレシーバーとして『SignUpというメソッド』を作る。
// 引数のuserにユーザーが入力したリクエストが入っている状態。
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	// パスワードをハッシュ化するための処理 10は暗号の複雑さを表す。
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	// CreateUserはユーザーオブジェうとのポインタを引数で受け取るので、&newUserでnewUserのポインタを取得し、引数で渡す。
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	// CreateUserに成功した場合、ポインタで渡したnewUserのオブジェクトの内容が、新しく作成したユーザーの内容に変わる。
	// そこからIDとEmailを取り出して、ユーザーのレスポンスの新しい構造体の実態を作成し、resUserという変数に格納してからreturnで返す。
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	// 成功している場合は、errorが発生しないので、nilになる。
	return resUser, nil
}

// ここでいったん区切れる。

// ログインメソッド
// userUsecaseをポインタレシーバーとして受け取る。
// 返り値はJWTのstringとerror
func (uu *userUsecase) Login(user model.User) (string, error) {

	// validation
	// 返り値の型が、stringとerrorになっているので、空の文字列とエラーを返す。
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	// ユーザーから送信されたEmailがデータベース内に存在するか判定する処理。
	// まず、Emailで間作するユーザーのオブジェクトを格納するための空のオブジェクトを作成。
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// 送られてきたEmailが存在する場合、パスワードの検証する。
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// パスワードが一致した場合、JWTトークンの生成を行う。
	// jwtパッケージのwithClaimsを使ってClaimsの設定を行う。
	// HS255というアルゴリズムを指定するのと、ペイロードの設定として、user_idとJWTの有効期限を設定している。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	// tokenで条件等を定義し、ここで実行して生成を行う。引数に、環境変数JWTのシークレットキーを設定。
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
