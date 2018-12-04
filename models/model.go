package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

/**
 * @Author: ruke
 * @Date: 2018-12-04 10:13:33
 * @Desc: 定义orm
 */
func SetOrm(Db *gorm.DB) {
	db = Db
}