package handler

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "../models"
    "fmt"
)

func CreatePost(c echo.Context) error {
    post := new(model.Post)
    if err := c.Bind(post); err != nil {
        return err
    }

    if post.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid to or message fields",
        }
    }

    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    post.UID = getUID(c)
    model.CreatePost(post)

    return c.Redirect(http.StatusFound, "/posts")
}

func CreateChildPost(c echo.Context) error {
    post := new(model.Post)
    if err := c.Bind(post); err != nil {
        return err
    }

    if post.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid to or message fields",
        }
    }

    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    parent_id, err := strconv.Atoi(c.Param("parent_id"))
    if err != nil {
        return echo.ErrNotFound
    }
    
    post.UID = getUID(c)
    post.ParentID = parent_id;

    model.CreatePost(post)

    return c.Redirect(http.StatusFound, "/posts/" + c.Param("parent_id"))
}

func GetPosts(c echo.Context) error {
    fmt.Println("hogehogeここまで！0")

    posts := model.GetParentPosts()

    fmt.Println("hogehogeここまで！１")
    // 一覧ページは誰でも見れるようにするので認証しない
    user := getUser(c)

    fmt.Println("hogehogeここまで！2")
    if user.ID == 0 {
        fmt.Println("hogehogeここまで！３")
        return c.Render(http.StatusOK, "post", posts)
    }
    fmt.Println("hogehogeここまで！４")

    return c.Render(http.StatusOK, "post", posts)
}

func GetPostDetails(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }
    
    parent_id, err := strconv.Atoi(c.Param("parent_id"))
    if err != nil {
        return echo.ErrNotFound
    }
    parent_post :=  model.FindPosts(&model.Post{ID: parent_id})
    child_posts := model.FindPosts(&model.Post{ParentID: parent_id})
    posts := model.ParentAndChildren{
        parent_post,
        child_posts,
    }

    return c.Render(http.StatusOK, "post_detail", posts)
}

func DeletePost(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)
    
    postID, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        return echo.ErrNotFound
    }

    if err := model.DeletePost(&model.Post{ID: postID, UID: uid}); err != nil {
        // TODO: 別の人の投稿だと分かるようにする
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    return c.Redirect(http.StatusFound, "/posts")
}

func UpdatePost(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)

    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.ErrNotFound
    }

    posts := model.FindPosts(&model.Post{ID: postID, UID: uid})
    if len(posts) == 0 {
        return echo.ErrNotFound
    }
    post := posts[0]
    post.Content = c.FormValue("content")
    if err := model.UpdatePost(&post); err != nil {
        return echo.ErrNotFound
    }

    return c.Redirect(http.StatusFound, "/posts/" + c.Param("id") + "/edit")
}

func PostUpdater(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)

    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.ErrNotFound
    }

    posts := model.FindPosts(&model.Post{ID: postID, UID: uid})
    if len(posts) == 0 {
        // TODO: 別の人の投稿だと分かるようにする
        return c.Render(http.StatusOK, "error_need_login", c)
    }
    post := posts[0]
    return c.Render(http.StatusOK, "post_updater", post)
}

func UserUpdater(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    return c.Render(http.StatusOK, "user_updater", user)
}

func UpdateUser(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    user.Name = c.FormValue("name")
    user.EMail = c.FormValue("email")

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

    if err := model.UpdateUser(&user); err != nil {
        // TODO アップデートに失敗したことが分かるようにする
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    return c.Redirect(http.StatusFound, "/user")
}

func DeleteUser(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)

    if err := model.DeleteUser(&model.User{ID: uid}); err != nil {
        return echo.ErrNotFound
    }

    r := c.Request()
    cookie, err := r.Cookie("uid")
    if err != nil {
        return echo.ErrNotFound 
    }

    // cookieを無効にする
    cookie.MaxAge = -1
    w := c.Response()
    http.SetCookie(w, cookie)

    return c.Redirect(http.StatusFound, "/login")
}

func Logout(c echo.Context) error {
    user := getUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    r := c.Request()
    cookie, err := r.Cookie("uid")
    if err != nil {
        return echo.ErrNotFound 
    }

    // cookieを無効にする
    cookie.MaxAge = -1
    w := c.Response()
    http.SetCookie(w, cookie)

    return c.Redirect(http.StatusFound, "/login")
}