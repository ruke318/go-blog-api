package main

import (
	// "fmt"
	"blog/routers"
	"blog/config"
	"blog/database"
	"blog/tools"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/olivere/elastic"
)

var (
	conf config.Config
	db *gorm.DB
)

func main () {
	// 加载配置文件
	conf = config.Load()
	db = database.ConnectMysql(conf.Mysql)
	db.SingularTable(true)
	//连接es
	esClient, err := elastic.NewClient(elastic.SetURL(conf.Elastic.Url))
	if err != nil {
		tools.Error("连接es失败")
	}
	// 加载路由, 将db也作为参数传递
	app := routers.Dispath(db, esClient)
	// 启动服务
	app.Run(iris.Addr(":8002"))
	//关闭数据库连接
	defer db.Close()
}