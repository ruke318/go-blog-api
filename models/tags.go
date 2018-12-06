package models

import "time"

type Tags struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	AddTime time.Time `gorm:"column:addTime;default:now()" json:"addTime"`
	AddTimeStr string `gorm:"-" json:"addTimeStr"`
}