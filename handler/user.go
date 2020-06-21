package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "../models"
)

// ユーザー情報編集画面を返す
func UserUpdater(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    return c.Render(http.StatusOK, "user_updater", user)
}

// ユーザー情報更新処理
func UpdateUser(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    user.Name = c.FormValue("name")
    user.EMail = c.FormValue("email")

    if user.Name == "" || user.EMail == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "名前かeMailに誤った値が使用されています",
        }
    }

    if u := model.FindUser(&model.User{EMail: user.EMail}); u.ID != 0 {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "そのメールアドレスは既に利用されています。",
        }
    }

    if err := model.UpdateUser(&user); err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "ユーザー情報の更新に失敗しました",
        }
    }

    return c.Redirect(http.StatusFound, "/user")
}
