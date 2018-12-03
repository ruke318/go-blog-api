package controllers

import (
	// "fmt"
	// "time"
	"blog/models"
	"blog/tools"
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
)

type UserCtr struct {}

var (
	userModel *models.Users
)

func (ctr *UserCtr) SetOrm(Db *gorm.DB) {
	models.SetOrm(Db)
	userModel = models.UserModel()
}

func (ctr *UserCtr) GetAll(request iris.Context) {
	ret := userModel.GetAll()
	res := tools.Error(ret, "test", tools.UserNotFound)
	request.JSON(res)
}

func (ctr *UserCtr) GetUserById(request iris.Context) {
	id := request.Params().Get("id")
	info := userModel.GetUserById(id)
	request.JSON(info)
}

func (ctr *UserCtr) AddUser(request iris.Context) {
	user := &models.Users{}
	request.ReadJSON(user)
	newUser := userModel.Create(*user)
	request.JSON(newUser)
}

func UserController() *UserCtr {
	return new(UserCtr)
}