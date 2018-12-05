/*
 * @Author: ruke
 * @Date: yyyy-12-dd 11:41:38
 * @Desc: 
 */
package controllers

import (
	// "fmt"
	"blog/models"
	"blog/tools"
	"github.com/kataras/iris"
)

type PostsCtr struct {}

var (
	postsModel *models.Posts
)

/**
 * @Author: ruke
 * @Date: 2018-12-04 15:45:18
 * @Desc: 通过过滤获取列表
 */
func (ctr *PostsCtr) GetList(request iris.Context) {
	navId := tools.ParseInt(request.FormValue("navId"), 0)
	author := tools.ParseInt(request.FormValue("author"), 0)
	// search := request.FormValue("search") //等待es介入
	tagId := tools.ParseInt(request.FormValue("tagId"), 0)
	page := tools.ParseInt(request.FormValue("page"), 1)
	pageSize := tools.ParseInt(request.FormValue("pageSize"), tools.DefaultPageSize)
	ret, current, count, hasNext := postsModel.GetAll(&models.PostFilter{NavId: navId, Author: author, TagId: tagId, Page: page, PageSize: pageSize})
	data := PageData{Data: ret, Count: count, HasNext: hasNext, Current: current}
	request.JSON(tools.Success(data))
}

func PostsController() *PostsCtr {
	return new(PostsCtr)
}