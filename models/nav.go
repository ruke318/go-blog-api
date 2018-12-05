package models

import (
	// "fmt"
	"github.com/thoas/go-funk"
)

type Nav struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Pid uint `json:"pid"`
	Icon string `gorm:"default:null" json:"icon"`
	Sub interface{} `gorm:"-" json:"sub"`
}

/**
 * @Author: ruke
 * @Date: 2018-12-05 15:43:21
 * @Desc: 获取nav
 */
func (model *Nav)GetNav() (navs []*Nav) {
	reDb := db
	//查询顶级目录
	db.Where("pid = ?", 0).Order("id").Find(&navs)
	//查询耳机目录
	all := []*Nav{}
	reDb.Where("pid <> ?", 0).Find(&all)
	//通过filter得到每个顶级目录中的子集
	for _, value := range navs {
		value.Sub = funk.Filter(all, func (item *Nav) bool{
			return value.ID == item.Pid
		})
	}
	return
}

/**
 * @Author: ruke
 * @Date: 2018-12-05 15:46:09
 * @Desc: 添加导航
 */
func (model *Nav) Create(nav Nav) Nav {
	db.Create(&nav)
	return nav
}