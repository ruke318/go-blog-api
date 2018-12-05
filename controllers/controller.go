package controllers

import (
	"blog/tools"
	"blog/models"
	"github.com/jinzhu/gorm"
)

// 订单返回参数
var Response = &tools.ApiResponse{}

type PageData struct{
	Data interface{} `json:"data"`
	Current int `json:"current"`
	Count int `json:"count"`
	HasNext bool `json:"hasNext"`
}

func SetOrm(Db *gorm.DB) {
	models.SetOrm(Db)
}