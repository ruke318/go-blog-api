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
	AddTime time.Time `gorm:"column:addTime;default:now()" json:"addTime"`
	Status int `gorm:"type:tinyint;default:1" json:"status"`
	Path string `json:"path"`
	AuthorInfo Users     `gorm:"ForeignKey:UserId;AssociationForeignKey:ID" json:"authorInfo"`
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
	newDb.Where("`key` = ?", key).Where("pid = ?", 0).Order("addTime desc").Offset(offset).Limit(pageSize).Preload("AuthorInfo").Find(&list)
	for _, item := range list {
		itemDb := db
		id := strconv.Itoa(int(item.ID))
		itemDb.Where("`key` = ?", key).Where("id <> ?", item.ID).Where("path like ? or path like ?", "%," + id, "%,"+ id +",%").Order("path").Preload("AuthorInfo").Find(&item.Sub)
	}
	return list, count, page, hasNext
}

/**
 * @Author: ruke
 * @Date: 2018-12-21 14:32:14
 * @Desc: 添加评论
 */
func (replys *Replys) Create(reply Replys) Replys {
	reply.Level = 0
	top := Replys{}
	if reply.Pid != 0 {
		db.First(&top, reply.Pid)
		reply.Level = top.Level + 1
	}
	db.Create(&reply)
	if reply.ID != 0 {
		if reply.Pid != 0 {
			reply.Path = top.Path + "," + strconv.Itoa(int(reply.ID))
		} else {
			reply.Path = "0," + strconv.Itoa(int(reply.ID))
		}
		db.Save(&reply)
	}
	return reply
}
