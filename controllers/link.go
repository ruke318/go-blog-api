/*
 * @Author: ruke
 * @Date: yyyy-12-dd 15:01:45
 * @Desc: 
 */
package controllers

import (
	"blog/models"
	"github.com/kataras/iris"
)

type LinkCtr struct {}

var (
	linkModel *models.Links
)

/**
 * @Author: ruke
 * @Date: 2018-12-03 15:10:09
 * @Desc: 获取所有链接 
 */
func (ctr *LinkCtr) GetAll(request iris.Context) {
	ret := linkModel.GetAll()
	request.JSON(ret)
}

func LinkController() *LinkCtr {
	return new(LinkCtr)
}