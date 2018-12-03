/*
 * @Author: ruke
 * @Date: yyyy-12-dd 15:01:45
 * @Desc: 
 */
package controllers

import (
	// "fmt"
	"blog/models"
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
)

type LinkCtr struct {}

var (
	linkModel *models.Links
)

/**
 * @Author: ruke
 * @Date: 2018-12-03 15:09:49
 * @Desc: 设置orm
 */
func (ctr *LinkCtr) SetOrm(Db *gorm.DB) {
	models.SetOrm(Db)
	linkModel = models.LinkModel()
}

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