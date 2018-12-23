/*
 * @Author: ruke
 * @Date: yyyy-12-dd 15:01:45
 * @Desc: 
 */
package controllers

import (
	"blog/models"
	"blog/tools"
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

func (ctr *LinkCtr) AddLink(request iris.Context) {
	link := models.Links{}
	request.ReadJSON(&link)
	newLink := linkModel.Create(link)
	if newLink.ID == 0 {
		Response = tools.Error("添加失败")
	} else {
		Response = tools.Success(newLink)
	}
	request.JSON(Response)
}

func LinkController() *LinkCtr {
	return new(LinkCtr)
}