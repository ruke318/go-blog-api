package main

import (
	// "fmt"
	"blog/config"
	"blog/database"
	"blog/routers"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

var (
	conf config.Config
	db   *gorm.DB
)

func main() {
	// 加载配置文件
	conf = config.Load()
	db = database.ConnectMysql(conf.Mysql)
	db.SingularTable(true)
	//连接es
	esClient := database.ConectElastic(conf.Elastic)
	// 加载路由, 将db也作为参数传递
	app := routers.Dispath(db, esClient)
	// 启动服务
	app.Run(iris.Addr(":8002"))
	//关闭数据库连接
	defer db.Close()
}
