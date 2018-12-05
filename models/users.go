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
	AddTime *time.Time `gorm:"column:addTime" json:"addTime"`
	Wxopenid string `json:"wxopenid"`
	UserResource string `gorm:"column:userResource;default:qq" json:"userResource"`
	Url string `json:"url"`
}

/**
 * @Author: ruke
 * @Date: 2018-12-03 16:30:36
 * @Desc: 获取所有用户
 */
func (user *Users) GetAll() (users []*Users) {
	db.Find(&users)
	return
}

/**
 * @Author: ruke
 * @Date: 2018-12-03 16:30:49
 * @Desc: 根据ID获取用户信息
 */
func (user *Users) GetUserById(id string) Users {
	users := Users{}
	db.First(&users, id)
	return users
}

func (user *Users) Create(users Users) Users {
	t := time.Now().Local()
	users.AddTime = &t
	db.Create(&users)
	return users
}