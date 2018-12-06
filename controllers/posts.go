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
	"strconv"
	"strings"
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

/**
 * @Author: ruke
 * @Date: 2018-12-06 15:38:33
 * @Desc: 添加文章详情
 */
func (ctr *PostsCtr) AddPosts(request iris.Context) {
	posts := models.Posts{}
	request.ReadJSON(&posts)
	newPosts := postsModel.Create(posts)
	if newPosts.ID == 0 {
		Response = tools.Error("添加失败")
	} else {
		Response = tools.Success(newPosts)
	}
	request.JSON(Response)
}

/**
 * @Author: ruke
 * @Date: 2018-12-06 15:38:50
 * @Desc: 查询文章详情
 */
func (ctr *PostsCtr) GetDetail(request iris.Context) {
	id, _ := strconv.Atoi(request.Params().Get("id"))
	detail := postsModel.Detail(id)
	if detail.ID == 0 {
		Response = tools.Error("文章不存在")
	} else {
		Response = tools.Success(detail)
	}
	request.JSON(Response)
}

/**
 * @Author: ruke
 * @Date: 2018-12-06 15:37:33
 * @Desc: 搜索文章
 */
func (ctr *PostsCtr) Search(request iris.Context) {
	keyword := request.FormValue("keyword")
	keyword = strings.Trim(keyword, "")
	page := tools.ParseInt(request.FormValue("page"), 1)
	pageSize := tools.ParseInt(request.FormValue("pageSize"), tools.DefaultPageSize)
	list, total, hasNext, current := postsModel.Search(keyword, page, pageSize)
	data := PageData{Data: list, Count: int(total), HasNext: hasNext, Current: current}
	request.JSON(tools.Success(data))
}

func PostsController() *PostsCtr {
	return new(PostsCtr)
}