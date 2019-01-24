package controllers

import (
	"blog/models"
	"blog/tools"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic"
)

// 订单返回参数
var Response = &tools.ApiResponse{}

type PageData struct {
	Data    interface{} `json:"data"`
	Current int         `json:"current"`
	Count   int         `json:"count"`
	HasNext bool        `json:"hasNext"`
}

func SetOrm(Db *gorm.DB, EsClient *elastic.Client) {
	models.SetOrm(Db, EsClient)
}
