package models

import (
	// "fmt"
	"time"
	// "github.com/jinzhu/gorm"
)

type Links struct {
	ID uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Desc string `json:"desc"`
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`
	Status uint `json:"status"`
	Logo string `json:"logo"`
}

func (user *Links) GetAll() Links {
    links := Links{}
	db.Find(&links)
	return links
}