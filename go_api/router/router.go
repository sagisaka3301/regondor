package router

import (
	"go_api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ルーターの中でタスクコントローラーを使用できるようにするために、引数にタスクコントローラーも追加。
func NewRouter(uc controller.IUserController, tc controller.ITaskController, mc controller.IMypageController) *echo.Echo {
	// echoのインスタンスに対し、エンドポイントを作成。
	e := echo.New()

	// CORSのミドルウェアを追加。
	// e.UseでCORSのミドルウェアを追加し、AllowOriginsのところにアクセスを許可するフロントエンドの列にドメインを追加。
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Reactで作ったアプリをバーセルにデプロイしたときにできるドメインをFR_URLに設定している。
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		// 許可するヘッダーを書いていく。echoのHeaderXCSRFTOKENを含めることによって、ヘッダー経由でCSRFトークンを受け取れるようにしている。
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		// 許可したいメソッドを追加。
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		// cookieの送受信を可能にするため、trueにする。
		AllowCredentials: true,
	}))
	// CSRFのミドルウェアを追加。CSRFトークンを格納するcookieのせってを行う。
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// cookieのパスとして、"/"(インデックス)
		// cookieのドメインとしてAPI_DOMAINを設定する。
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// NoneModeにしてしまうと、自動的にsecureモードがfalseになり、postmanで動作確認ができなくなる。よって、defaultModeにしておく。
		CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		// 有効期限：デフォルトは24時間(秒単位なので、この場合60秒)
		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp) // signupのエンドポイントにリクエストがあったときは、コントローラーのSignUpメソッドを呼び出す。
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	// taskについて。前回作成したインスタンスのeに対し、新しくグループを作る。
	// エンドポイントをグループ化してtという変数に格納
	// そして、タスクのグループに対し、JWTのミドルウェアを適用するようにする。(middleware:authと同じ。)
	t := e.Group("/tasks")
	// Useキーワードを使うことで、エンドポイントにミドルウェアを追加することができる。
	// echoのjwtというミドルウェアを適用している。
	t.Use(echojwt.WithConfig(echojwt.Config{
		// jwtを生成したときと同じシークレットキーを指定する。
		SigningKey: []byte(os.Getenv("SECRET")),
		// クライアントから送られてくるjwtトークンがどこに格納されているのか指定する必要がある。
		// 今回はcookieの中にtokenという名前でjwtトークンを格納するように実装しているのでこの書き方。
		TokenLookup: "cookie:token",
	}))
	// タスク関連のエンドポイントを追加しておく。
	// グループ化されているので、xxx.com/tasks/以降のurlになる。
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	m := e.Group("/mypage")
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	m.GET("", mc.GetUser)

	return e
}
