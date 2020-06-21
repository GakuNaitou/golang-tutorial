package model

import "fmt"

type Post struct {
    UID       int    `json:"uid"`
    ID        int    `json:"id" gorm:"praimaly_key"`
    Content      string `json:"content"`
    ParentID    int    `json:"parent_id"`
}

type Posts []Post

type ParentAndChildren struct {
    ParentPost Posts
    ChildPosts Posts
}

func CreatePost(post *Post) {
    db.Create(post)
}

func FindPosts(t *Post) Posts {
    var posts Posts
    db.Where(t).Find(&posts)
    return posts
}

func GetParentPosts() Posts {
    var parentPosts Posts
    db.Where("parent_id = ?", 0).Find(&parentPosts)
    return parentPosts
}

func DeletePost(t *Post) error {
    if rows := db.Where(t).Delete(&Post{}).RowsAffected; rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to delete", t)
    }
    return nil
}

func UpdatePost(t *Post) error {
    rows := db.Model(t).Update(map[string]interface{}{
        "content": t.Content,
    }).RowsAffected
    if rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to update", t)
    }
    return nil
}
