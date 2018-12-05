package models

import (
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic"
)

var (
	db *gorm.DB
	esClient *elastic.Client
)

/**
 * @Author: ruke
 * @Date: 2018-12-04 10:13:33
 * @Desc: 定义orm
 */
func SetOrm(Db *gorm.DB, EsClient *elastic.Client) {
	db = Db
	esClient = EsClient
}