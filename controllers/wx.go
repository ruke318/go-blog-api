package controllers

import (
	"github.com/kataras/iris"
	"blog/services"
	"blog/tools"
)

var (
	wxService *services.WxApp
)

type WxCtr struct {}

func WxController() *WxCtr {
	return new(WxCtr)
}

func (ctr *WxCtr)Login(request iris.Context) {
	code := request.Params().Get("code")
	str := wxService.Login(code)
	Response = tools.Success(str)
	request.JSON(Response)
}