package controllers

import (
	// "fmt"
	// "time"
	"blog/models"
	"blog/tools"
	"github.com/kataras/iris"
)

type UserCtr struct {}

var (
	userModel *models.Users
)

/**
 * @Author: ruke
 * @Date: 2018-12-04 10:14:58
 * @Desc: 获取所有用户
 */
func (ctr *UserCtr) GetAll(request iris.Context) {
	ret := userModel.GetAll()
	res := tools.Success(ret)
	request.JSON(res)
}

/**
 * @Author: ruke
 * @Date: 2018-12-04 10:15:24
 * @Desc: 根据ID获取用户信息
 */
func (ctr *UserCtr) GetUserById(request iris.Context) {
	id := request.Params().Get("id")
	info := userModel.GetUserById(id)
	if info.ID == 0 {
		Response = tools.Error("用户不存在", tools.UserNotFound)
	} else {
		Response = tools.Success(info)
	}
	request.JSON(Response)
}

func (ctr *UserCtr) AddUser(request iris.Context) {
	user := &models.Users{}
	request.ReadJSON(user)
	newUser := userModel.Create(*user)
	if newUser.ID == 0 {
		Response = tools.Error("添加用户失败", tools.UserAddFail)
	} else {
		Response = tools.Success(newUser)
	}
	request.JSON(Response)
}

func UserController() *UserCtr {
	return new(UserCtr)
}