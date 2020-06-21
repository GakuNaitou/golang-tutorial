package main
import (
	"html/template"
	"io"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
	"./models" // DBの定義と処理
	"./handler" // アプリケーションサーバーの処理
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// TODO: カラム名の変更、削除をできるようにする
	model.InitDB() // DBの初期化処理(マイグレーション実行)

	router := newRouter()
	router.Logger.Fatal(router.Start(":1323"))
}

// routing
func newRouter() *echo.Echo {

	// views/以下のhtmlファイルをtemplateとして使用できるように定義する
	// https://echo.labstack.com/guide/templates
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	// リクエスト単位のログを、起動したhttpサーバー上に出力する
	e.Use(middleware.Logger())

	// 認証周り
	e.GET("/login", handler.LoginPage)
	e.GET("/signup", handler.UserRegister)
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.GET("/logout", handler.Logout) 
	e.POST("/user/delete", handler.DeleteAccount) 

	// ユーザー周り
	e.GET("/user", handler.UserUpdater) 
	e.POST("/user", handler.UpdateUser)

	// 投稿周り
	e.GET("/posts", handler.GetPosts)
	e.POST("posts", handler.CreatePost) 
	e.GET("posts/:parent_id", handler.GetPostDetails)
	e.POST("posts/:parent_id", handler.CreateChildPost)
	e.POST("/posts/delete/:id", handler.DeletePost)
	e.GET("/posts/:id/edit", handler.PostUpdater) 
	e.POST("/posts/:id/edit", handler.UpdatePost) 
	
	return e
}
