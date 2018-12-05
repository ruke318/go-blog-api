package models

import "time"

type Tags struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	AddTime time.Time `grom:"column:addTime" json:"addTime"`
}