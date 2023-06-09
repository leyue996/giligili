package app

import (
	"giligili/app/config"
	"giligili/app/http/dao"
	"strings"
)

var Config config.Conf

func Init() {

	Config.ConfInit() //本地环境变量

	path := strings.Join([]string{Config.Mysql.DbUser, ":", Config.Mysql.DbPassWord, "@tcp(", Config.Mysql.DbHost, ":", Config.Mysql.DbPort, ")/", Config.Mysql.DbName, "?charset=utf8&parseTime=true"}, "")

	dao.Database(path) // 连接数据库
}
