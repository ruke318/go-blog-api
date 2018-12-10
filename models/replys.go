package models

import (
	"time"
	"strconv"
)

type Replys struct {
	ID uint `gorm:"primary_key" json:"id"`
	Pid int `json:"pid"`
	Level int `json:"level"`
	Content string `gorm:"type:text" json:"content"`
	UserId int `gorm:"column:userId" json:"userId"`
	Key string `json:"key"`
	Ip string `json:"ip"`
	Os string `json:"os"`
	Tool string `json:"tool"`
	AddTime time.Time `gorm:"column:addTime" json:"addTime"`
	Status int `gorm:"type:tinyint" json:"status"`
	Path string `json:"path"`
	Sub []*Replys `json:"sub" grom:"-"`
}

/**
 * @Author: ruke
 * @Date: 2018-12-10 16:39:22
 * @Desc: 获取评论列表
 */
func (replys *Replys)GetList(key string, page int, pageSize int) ([]*Replys, int, int, bool) {
	offset := (page - 1) * pageSize
	findDb := db
	list := []*Replys{}
	count := 0
	newDb := db
	findDb.Where("`key` = ?", key).Where("pid = ?", 0).Find(&list).Count(&count)
	hasNext := page * pageSize < count
	newDb.Where("`key` = ?", key).Where("pid = ?", 0).Order("addTime desc").Offset(offset).Limit(pageSize).Find(&list)
	for _, item := range list {
		itemDb := db
		id := strconv.Itoa(int(item.ID))
		itemDb.Where("`key` = ?", key).Where("id <> ?", item.ID).Where("path like ? or path like ?", "%," + id, "%,"+ id +",%").Order("path").Find(&item.Sub)
	}
	return list, count, page, hasNext
}