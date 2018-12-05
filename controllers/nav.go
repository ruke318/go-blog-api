package controllers

import (
	"blog/models"
	"blog/tools"
	"github.com/kataras/iris"
)

type NavCtr struct {}

var (
	navModel *models.Nav
)

/**
 * @Author: ruke
 * @Date: 2018-12-05 15:46:35
 * @Desc: 获取所有导航
 */
func (ctr *NavCtr)GetNav(request iris.Context) {
	navs := navModel.GetNav()
	Response = tools.Success(navs)
	request.JSON(Response)
}

/**
 * @Author: ruke
 * @Date: 2018-12-05 15:52:07
 * @Desc: 创建nav
 */
func (ctr *NavCtr) AddNav(request iris.Context) {
	add := models.Nav{}
	request.ReadJSON(&add)
	newAdd := navModel.Create(add)
	if newAdd.ID == 0 {
		Response = tools.Error("添加失败")
	} else {
		Response = tools.Success(newAdd)
	}
	request.JSON(Response)
}

func NavController() *NavCtr {
	return new(NavCtr)
}