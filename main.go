package main
import (
	"html/template"
	"io"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"./models" // DBの定義と処理
	"./handler"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	model.InitDB() // DBの初期化処理(マイグレーション実行)

	router := newRouter()
	router.Logger.Fatal(router.Start(":1323"))
}

// routing
func newRouter() *echo.Echo {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())

	e.GET("/login", Login)
	e.GET("/signup", UserRegister)
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	e.GET("/user", handler.UserUpdater) 
	e.POST("/user", handler.UpdateUser) 
	e.GET("/logout", handler.Logout) 

	e.GET("/posts", handler.GetPosts)
	e.POST("posts", handler.CreatePost) 
	e.GET("posts/:parent_id", handler.GetPostDetails)
	e.POST("posts/:parent_id", handler.CreateChildPost)
	e.POST("/posts/delete/:id", handler.DeletePost)
	e.GET("/posts/:id/edit", handler.PostUpdater) 
	e.POST("/posts/:id/edit", handler.UpdatePost) 
	
	return e
}

func Login(c echo.Context) error {
	db := DBConnect()
	return c.Render(http.StatusOK, "login", db)
}

func UserRegister(c echo.Context) error {
	db := DBConnect()
	return c.Render(http.StatusOK, "signup", db)
}

// DB接続
func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "docker_user"
	PASS := "docker_user_pwd"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "docker_db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
			panic(err.Error())
	}
	return db
}