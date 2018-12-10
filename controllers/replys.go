package controllers

import (
	"blog/models"
	"blog/tools"
	"github.com/kataras/iris"
)

type ReplysCtr struct{}

var (
	replysModel *models.Replys
)

/**
 * @Author: ruke
 * @Date: 2018-12-10 16:39:44
 * @Desc: 获取评论列表
 */
func (ctr *ReplysCtr) GetList(request iris.Context) {
	key := request.FormValue("key")
	page := tools.ParseInt(request.FormValue("page"), 1)
	pageSize := tools.ParseInt(request.FormValue("pageSize"), tools.DefaultPageSize)
	ret, total, current, hasNext := replysModel.GetList(key, page, pageSize)
	data := PageData{Data: ret, Count: total, HasNext: hasNext, Current: current}
	Response = tools.Success(data)
	request.JSON(Response)
}

func ReplysController() *ReplysCtr {
	return new(ReplysCtr)
}