package model

import (
    "fmt"
    "time"
)

type Post struct {
    ID        int    `json:"id" gorm:"praimaly_key"`
    UID       int    `json:"uid"`
    Content      string `json:"content"`
    ParentID    int    `json:"parent_id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

type Posts []Post

type ParentAndChildren struct {
    ParentPost PostAndUserName
    ChildPosts []PostAndUserName
}

type PostAndUserName struct {
    Post Post
    UserName string
}

func CreatePost(post *Post) {
    db = DBConnect()
    db.Create(post)
}

func FindPosts(t *Post) Posts {
    db = DBConnect()
    var posts Posts
    db.Order("created_at desc").Where(t).Find(&posts)
    return posts
}

func GetParentPosts() Posts {
    db = DBConnect()
    var parentPosts Posts
    db.Order("created_at desc").Where("parent_id = ?", 0).Find(&parentPosts)
    return parentPosts
}

func DeletePost(t *Post) error {
    db = DBConnect()
    if rows := db.Where(t).Delete(&Post{}).RowsAffected; rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to delete", t)
    }
    return nil
}

func UpdatePost(t *Post) error {
    db = DBConnect()
    rows := db.Model(t).Update(map[string]interface{}{
        "content": t.Content,
    }).RowsAffected
    if rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to update", t)
    }
    return nil
}
