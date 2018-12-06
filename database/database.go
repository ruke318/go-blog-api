/**
 * @Author: ruke
 * @Date: 2018-12-03 13:40:53
 * @Desc:  数据库连接方案
 */
package database

import (
	"fmt"

	"blog/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/olivere/elastic"
)

/**
 * @Author: ruke
 * @Date: 2018-12-03 13:41:17
 * @Desc: MySQL 连接
 */
func ConnectMysql(conf *config.Mysql) *gorm.DB {
	return Connect("mysql", conf.Connect)
}

/**
 * @Author: ruke
 * @Date: 2018-12-03 13:41:43
 * @Desc: redis连接
 */
func ConnectRedis(conf *config.Redis) *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password, // no password set
		DB:       int(conf.Db),   // 因为系统是64位的，所以默认的 int 型是 int64
	})

	return rd
}

/**
 * @Author: ruke
 * @Date: 2018-12-06 16:38:17
 * @Desc: 链接es
 */
func ConectElastic(conf *config.Elastic) *elastic.Client {
	esClient, err := elastic.NewClient(elastic.SetURL(conf.Url))
	if err != nil {
		panic(fmt.Sprintf("es 链接失败 err=%+v", err))
	}
	return esClient
}

func Connect(driver string, conf string) *gorm.DB {
	DB, err := gorm.Open(driver, conf)

	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}

	return DB
}