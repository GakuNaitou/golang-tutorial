package handler

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "../models"
    // "fmt"
)

// 投稿一覧画面を返す
func GetPosts(c echo.Context) error {
    posts := model.GetParentPosts()
    var post_data []model.PostAndUserName

    for _, post := range posts {
        user := model.FindUser(&model.User{ID: post.UID})
        post_and_user_name := model.PostAndUserName{
            post,
            user.Name,
        }
        post_data = append(post_data, post_and_user_name)
    }

    return c.Render(http.StatusOK, "post", post_data)
}

// 投稿詳細画面を返す
func GetPostDetails(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }
    
    parent_id, err := strconv.Atoi(c.Param("parent_id"))
    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "処理に失敗しました",
        }
    }
    
    parent_post :=  model.FindPosts(&model.Post{ID: parent_id})
    parent_user := model.FindUser(&model.User{ID: parent_post[0].UID})
    // var parent_post_data []model.PostAndUserName
    
    parent_post_data := model.PostAndUserName{
        parent_post[0],
        parent_user.Name,
    }

    child_posts := model.FindPosts(&model.Post{ParentID: parent_id})
    var child_post_data []model.PostAndUserName

    for _, post := range child_posts {
        user := model.FindUser(&model.User{ID: post.UID})
        post_and_user_name := model.PostAndUserName{
            post,
            user.Name,
        }
        child_post_data = append(child_post_data, post_and_user_name)
    }

    
    posts := model.ParentAndChildren{
        parent_post_data,
        child_post_data,
    }

    return c.Render(http.StatusOK, "post_detail", posts)
}

// 投稿編集画面を返す
func PostUpdater(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)

    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "処理に失敗しました",
        }
    }

    posts := model.FindPosts(&model.Post{ID: postID})
    if len(posts) == 0 {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "投稿を見つけられませんでした",
        }
    }

    post := posts[0]

    if post.UID != uid {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "別の人の投稿は変更できません",
        }
    }

    return c.Render(http.StatusOK, "post_updater", post)
}

// 投稿作成処理
func CreatePost(c echo.Context) error {
    post := new(model.Post)
    if err := c.Bind(post); err != nil {
        return err
    }

    if post.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "投稿内容を入力してください",
        }
    }

    user := getCurrentUser(c)

    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    post.UID = getUID(c)
    model.CreatePost(post)

    return c.Redirect(http.StatusFound, "/posts")
}

// 子投稿作成処理
func CreateChildPost(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    post := new(model.Post)
    if err := c.Bind(post); err != nil {
        return err
    }

    if post.Content == "" {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "投稿内容を入力してください",
        }
    }

    parent_id, err := strconv.Atoi(c.Param("parent_id"))
    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "処理に失敗しました",
        }
    }
    
    post.UID = getUID(c)
    post.ParentID = parent_id;

    model.CreatePost(post)

    return c.Redirect(http.StatusFound, "/posts/" + c.Param("parent_id"))
}

// 投稿更新処理
func UpdatePost(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)

    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "処理に失敗しました",
        }
    }

    posts := model.FindPosts(&model.Post{ID: postID})
    if len(posts) == 0 {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "投稿を見つけられませんでした",
        }
    }
    
    post := posts[0]

    if post.UID != uid {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "別の人の投稿は変更できません",
        }
    }

    post.Content = c.FormValue("content")
    if err := model.UpdatePost(&post); err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "投稿の更新に失敗しました",
        }
    }

    return c.Redirect(http.StatusFound, "/posts/" + c.Param("id") + "/edit")
}

// 投稿削除処理
func DeletePost(c echo.Context) error {
    user := getCurrentUser(c)
    if user.ID == 0 {
        return c.Render(http.StatusOK, "error_need_login", c)
    }

    uid := getUID(c)
    
    postID, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "処理に失敗しました",
        }
    }

    if err := model.DeletePost(&model.Post{ID: postID, UID: uid}); err != nil {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "別の人の投稿は変更できません",
        }
    }

    return c.Redirect(http.StatusFound, "/posts")
}
