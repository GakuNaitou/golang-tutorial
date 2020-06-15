package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "../models"
    "strconv"
)

func Signup(c echo.Context) error {
    user := new(model.User)

    if err := c.Bind(user); err != nil {
        return err
    }

    if user.Name == "" || user.EMail == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid user_id or password",
        }
    }

    if u := model.FindUser(&model.User{EMail: user.EMail}); u.ID != 0 {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "email already exists",
        }
    }

    model.CreateUser(user)
    return c.Redirect(http.StatusFound, "/login")
}

func Login(c echo.Context) error {
    u := new(model.User)
    if err := c.Bind(u); err != nil {
        return err
    }

    user := model.FindUser(&model.User{EMail: u.EMail, Name: u.Name})
    if user.ID == 0 || user.EMail != u.EMail {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid name or email",
        }
    }

    userID := strconv.Itoa(user.ID)

    cookie := &http.Cookie{
        Name: "uid", // ここにcookieの名前を記述
        Value: userID, // ここにcookieの値を記述
        MaxAge: 1800,
    }

    w := c.Response()
    http.SetCookie(w, cookie)

    return c.Redirect(http.StatusFound, "/posts")
}

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
