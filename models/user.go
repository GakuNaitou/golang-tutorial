package model

import "fmt"

type User struct {
    ID       int   `json:"id" gorm:"praimaly_key"`
    Name     string `json:"name"`
    EMail string `json:"email"`
}

func CreateUser(user *User) {
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
    rows := db.Model(t).Update(map[string]interface{}{
        "name": t.Name,
        "email": t.EMail,
    }).RowsAffected
    if rows == 0 {
        return fmt.Errorf("Could not find Post (%v) to update", t)
    }
    return nil
}