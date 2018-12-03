package controllers

import (
	// "fmt"
	"blog/models"
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
	request.JSON(ret)
}

func UserController() *UserCtr {
	return new(UserCtr)
}