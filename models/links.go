package models

import (
	// "fmt"
	"time"
	// "github.com/jinzhu/gorm"
)

type Links struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Desc      string    `json:"desc"`
	Timestamp time.Time `gorm:"column:timestamp;default:now()" json:"timestamp"`
	Status    uint      `json:"status;default: 1"`
	Logo      string    `json:"logo"`
}

func (link *Links) GetAll() []Links {
	links := []Links{}
	db.Find(&links)
	return links
}

/**
 * @Author: ruke
 * @Date: 2018-12-23 16:12:40
 * @Desc: 添加链接
 */
func (model *Links) Create(link Links) Links {
	db.Create(&link)
	return link
}
