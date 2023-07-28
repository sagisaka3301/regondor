package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// コントローラーのインターフェースを作成。
// フレームワークのechoを使用しており、引数ではechoで定義されているContext型を指定。
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

// 構造体。フィールドとして、usecaseパッケージ内のIUserUsecaseインターフェース型の値をuuという名前で定義。
type userController struct {
	uu usecase.IUserUsecase
}

// controllerに対して、usecaseを依存性注入したいので、controller内にもコンストラクターを追加。
// 外側でインスタンス化されるusecaseを引数として注入できるようにしておく。
// 受け取ったusecaseのインスタンス(uu)を使ってuserControllerの実体を作る。そして、その値をアドレスを返す。
// 返り値の型はIUserController(インターフェース型)にしている。
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// userController型がIUserControllerインターフェースを満たすためにはすべてのメソッドを実装する。
func (uc *userController) SignUp(c echo.Context) error {
	// ユーザーから受け取るリクエスト(body)の値を構造体に変換する処理
	user := model.User{}
	// echoで用意されているBindメソッドを実行。引数にはユーザーオブジェクトのポインタ。
	// そうすることで、リクエストの値をユーザーオブジェクトのポインタが指し示す先の値(それは何？)に格納してくれる。
	if err := c.Bind(&user); err != nil {
		// 変換作業に失敗した場合、JSON形式でエラー文を返す。
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Bindに成功した場合、userusecaseのSignUpメソッドを実行。
	// コントローラーのSignUpとは別物なので注意。
	// 失敗した場合、InternalServerErrorを返す。
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 成功した場合、成功のステータスと新しく作成したユーザーの情報を返す。
	return c.JSON(http.StatusCreated, userRes)
}

// ログインメソッド
func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// 成功した場合は、取得したJWTトークンをサーバーサイドでクッキーに設定していきます。
	// まずはnew関数を使って。HTTPパッケージに定義されてるクッキー構造体を新しく作成します。
	cookie := new(http.Cookie)
	cookie.Name = "token"      // 任意の名前
	cookie.Value = tokenString //作成したJWTトークンを代入
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"                       // パスはindexを示す/
	cookie.Domain = os.Getenv("API_DOMAIN") // 環境変数のAPI_DOMAINを指定。
	cookie.Secure = true                    // postmanで動作確認したいので、コメントアウトでfalseにしておく。
	cookie.HttpOnly = true                  // trueにして、クライアントのJavaScriptからトークンの値を読み取れないようにする。
	cookie.SameSite = http.SameSiteNoneMode // FE BE違うドメイン(クロスドメイン)でのcookie送受信なので、SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// ログアウト
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""           // cookieの値を空にしたいので空の文字列に。
	cookie.Expires = time.Now() // 有効期限がすぐに切れるように指定。(これは初期化なので、一瞬で空のcookieを作るという流れ。)
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	// csrfトークンはechoのcontextの中で、csrfというキーワードで取得できる。
	// string型にアサーションしてからJSONでクライアントにcsrfトークンを返す。
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
