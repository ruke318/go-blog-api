package controllers

import (
	"blog/tools"
)

// 订单返回参数
var Response = &tools.ApiResponse{}

type PageData struct{
	Data interface{} `json:"data"`
	Current int `json:"current"`
	Count int `json:"count"`
	HasNext bool `json:"hasNext"`
}