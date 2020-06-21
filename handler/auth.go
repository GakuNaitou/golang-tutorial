package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "../models"
    "strconv"
)

// ユーザー登録画面を返す
func UserRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", nil)
}

// ログイン画面を返す
func LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

// ユーザー登録処理
func Signup(c echo.Context) error {
    user := new(model.User)
    if err := c.Bind(user); err != nil {
        return err
    }

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

    model.CreateUser(user)
    return c.Redirect(http.StatusFound, "/login")
}

// ログイン処理
func Login(c echo.Context) error {
    u := new(model.User)
    if err := c.Bind(u); err != nil {
        return err
    }

    user := model.FindUser(&model.User{EMail: u.EMail, Name: u.Name})
    if user.ID == 0 || user.EMail != u.EMail {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "名前かeMailに誤った値が使用されています",
        }
    }

    userID := strconv.Itoa(user.ID)

    cookie := &http.Cookie{
        Name: "uid", // cookieの名前
        Value: userID, // cookieの値
        MaxAge: 1800, // 有効期限(s)
    }

    w := c.Response()
    http.SetCookie(w, cookie)

    return c.Redirect(http.StatusFound, "/posts")
}

// ログアウト処理
func Logout(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    r := c.Request()
    cookie, err := r.Cookie("uid")
    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "処理に失敗しました",
        }
    }

    // cookieの有効期限をマイナス値にして無効にする
    cookie.MaxAge = -1
    w := c.Response()
    http.SetCookie(w, cookie)

    return c.Redirect(http.StatusFound, "/login")
}

// 退会処理
func DeleteAccount(c echo.Context) error {
    user := getCurrentUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)

    if err := model.DeleteUser(&model.User{ID: uid}); err != nil {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "ユーザーの削除に失敗しました",
        }
    }

    r := c.Request()
    cookie, err := r.Cookie("uid")
    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "処理に失敗しました",
        }
    }

    // cookieの有効期限をマイナス値にして無効にする
    cookie.MaxAge = -1
    w := c.Response()
    http.SetCookie(w, cookie)

    return c.Redirect(http.StatusFound, "/login")
}

// 現在のユーザーを返す
func getCurrentUser(c echo.Context) model.User {
    uid := getUID(c)
    user := model.FindUser(&model.User{ID: uid})
    return user
}

// 現在のユーザーのIDを返す
func getUID(c echo.Context) int {
    r := c.Request()
    cookie, err := r.Cookie("uid")

    if err != nil || cookie.Value == "0" {
        return -1
    }

    uid, err := strconv.Atoi(cookie.Value)
    if err != nil {
        return -1
    }

    return uid
}
