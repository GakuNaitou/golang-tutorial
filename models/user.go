package model

import (
    "fmt"
    "time"
)

type User struct {
    ID       int   `json:"id" gorm:"praimaly_key"`
    Name     string `json:"name"`
    EMail string `json:"e_mail"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

func CreateUser(user *User) {
    db = DBConnect()
    db.Create(user)
}

func FindUser(u *User) User {
    db = DBConnect()
    var user User
    db.Where(u).First(&user)
    fmt.Println(user)
    return user
}

func UpdateUser(t *User) error {
    db = DBConnect()
    rows := db.Model(t).Update(map[string]interface{}{
        "name": t.Name,
        "e_mail": t.EMail,
    }).RowsAffected
    if rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to update", t)
    }
    return nil
}

func DeleteUser(t *User) error {
    db = DBConnect()
    if rows := db.Where(t).Delete(&User{}).RowsAffected; rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to delete", t)
    }
    return nil
}