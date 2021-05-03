package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func InsertUser(user *User) {
	Db.NewRecord(user)
	Db.Create(&user)
}

func GetAllUsers(users *[]User) {
	Db.Find(&users)
}
