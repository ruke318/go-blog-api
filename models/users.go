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

/**
 * @Author: ruke
 * @Date: 2018-12-03 16:30:36
 * @Desc: 获取所有用户
 */
func (user *Users) GetAll() Users {
    users := Users{}
	db.Find(&users)
	return users
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

func UserModel() *Users {
	return new(Users)
}