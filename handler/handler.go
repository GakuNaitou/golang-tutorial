package handler

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "../models"
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

    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
    }

    post.UID = uid
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

    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
    }

    parent_id, err := strconv.Atoi(c.Param("parent_id"))
    if err != nil {
        return echo.ErrNotFound
    }

    post.UID = uid;
    post.ParentID = parent_id;

    model.CreatePost(post)

    return c.Redirect(http.StatusFound, "/posts/" + c.Param("parent_id"))
}

func GetPosts(c echo.Context) error {
    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
    }

    posts := model.GetParentPosts()
    return c.Render(http.StatusOK, "post", posts)
}

func GetPostDetails(c echo.Context) error {
    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
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
    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
    }
    
    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.ErrNotFound
    }

    if err := model.DeletePost(&model.Post{ID: postID, UID: uid}); err != nil {
        return echo.ErrNotFound
    }

    return c.Redirect(http.StatusFound, "/posts")
}

func UpdatePost(c echo.Context) error {
    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
    }

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
    uid := getUID(c)
    if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
    }

    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.ErrNotFound
    }

    posts := model.FindPosts(&model.Post{ID: postID, UID: uid})
    post := posts[0]

    return c.Render(http.StatusOK, "post_updater", post)
}

func UserUpdater(c echo.Context) error {
    uid := getUID(c)
    user := model.FindUser(&model.User{ID: uid})
    if user.ID == 0 {
        return echo.ErrNotFound
    }

    return c.Render(http.StatusOK, "user_updater", user)
}

func UpdateUser(c echo.Context) error {
    uid := getUID(c)
    user := model.FindUser(&model.User{ID: uid})
    if user.ID == 0 {
        return echo.ErrNotFound
    }

    user.Name = c.FormValue("name")
    user.EMail = c.FormValue("email")

    if err := model.UpdateUser(&user); err != nil {
        return echo.ErrNotFound
    }

    return c.Render(http.StatusOK, "user_updater", user)
}

func Logout(c echo.Context) error {
    uid := getUID(c)
    user := model.FindUser(&model.User{ID: uid})
    if user.ID == 0 {
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