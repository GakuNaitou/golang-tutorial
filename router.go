package main
import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
)

func newRouter() *echo.Echo {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())

	e.GET("/login", Login)
	e.GET("/user_regist", UserRegister)
	e.POST("/user_regist", RegistUser)
	e.POST("/login", RegistUser)
	
	return e
}