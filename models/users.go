package models

import (
	// "fmt"
	"time"
)

type Users struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Openid string `json:"openid"`
	AddTime time.Time `gorm:"column:addTime" json:"addTime"`
	Wxopenid string `json:"wxopenid"`
	UserResource string `gorm:"column:userResource" json:"userResource"`
}

func (user *Users) GetAll() Users {
    users := Users{}
	db.Find(&users)
	return users
}

func (user *Users) GetUserById(id int) Users {
	users := Users{}
	// db.First(&users, id)
	return users
}

func UserModel() *Users {
	return new(Users)
}