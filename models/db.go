package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func InitDB() {
	db := DBConnect()
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "docker_user"
	PASS := "docker_user_pwd"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "docker_db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
			panic(err.Error())
	}
	return db
}