
package config

import (
	"github.com/jinzhu/configor"
)

//mysql 配置
type Mysql struct {
	Connect     string	`json:"connect"`
}

// redis配置
type Redis struct {
	Host  string	`json:"host"`
	Password string `json:"password"`
	Db	int			`json:"db"`
}

type Elastic struct {
	Url string `json:"url"`
}

//总配置 
type Config struct {
	Mysql *Mysql
	Redis *Redis
	Elastic *Elastic
}

/**
 * @Author: ruke
 * @Date: 2018-12-03 13:39:46
 * @Desc: 加载配置
 */
func Load() Config {
	conf := Config{}
	configor.Load(&conf, "./config/config.json")
	return conf
}