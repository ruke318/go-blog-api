package models

import (
	"time"
	// "fmt"
)

type Posts struct {
	ID uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Desc string `gorm:"type:text" json:"desc"`
	AddTime time.Time `gorm:"column:addTime" json:"addTime"`
	Author int `json:"Author"`
	Tags string `json:"tags"`
	Imgs string `json:"imgs"`
	View uint `json:"view"`
	NavId int `gorm:"column:navId" json:"navId"`
}

/**
 * @Author: ruke
 * @Date: 2018-12-04 15:46:40
 * @Desc: 文章过滤结构
 */
type PostFilter struct {
	NavId int
	TagId int
	Author int
	Page int
	PageSize int
	Title string
}

func PostsModel() *Posts {
	return new(Posts)
}

/**
 * @Author: ruke
 * @Date: 2018-12-04 15:45:51
 * @Desc: 获取所有列表包括过滤分页
 */
func (post *Posts) GetAll(filter *PostFilter) ([]*Posts, int, int, bool) {
	posts := []*Posts{}
	findPost := &Posts{}
	findDb := db
	if filter.NavId != 0 {
		findPost.NavId = filter.NavId
	}

	if filter.Author != 0 {
		findPost.Author = filter.Author
	}

	if filter.TagId != 0 {
		findDb = findDb.Where("find_in_set(?, tags)", filter.TagId)
	}
	count := 0
	current := filter.Page
	hasNext := false
	findDb.Where(findPost).Find(&posts).Count(&count)
	if filter.Page > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		if filter.Page * filter.PageSize < count {
			hasNext = true
		}
		findDb.Where(findPost).Order("addTime desc").Offset(offset).Limit(filter.PageSize).Find(&posts)
	}
	return posts, current, count, hasNext
}