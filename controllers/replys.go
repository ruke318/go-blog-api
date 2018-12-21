package controllers

import (
	// "fmt"
	"blog/models"
	"blog/tools"
	"github.com/kataras/iris"
	"github.com/mssola/user_agent"
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

/**
 * @Author: ruke
 * @Date: 2018-12-21 14:32:39
 * @Desc: 添加评论
 */
func (ctr *ReplysCtr) AddReply(request iris.Context) {
	reply := models.Replys{}
	request.ReadJSON(&reply)
	userAgent := request.GetHeader("user-agent")
	ua := user_agent.New(userAgent)
	name, version := ua.Browser()
	reply.Os = ua.OS()
	reply.Tool = name + " " + version
	reply.Ip = request.RemoteAddr()
	newReply := replysModel.Create(reply)
	if newReply.ID == 0 {
		Response = tools.Error("添加失败")
	} else {
		Response = tools.Success(newReply)
	}
}

func ReplysController() *ReplysCtr {
	return new(ReplysCtr)
}