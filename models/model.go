package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetOrm(Db *gorm.DB) {
	db = Db
}