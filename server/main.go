package main

import (
	"wiliwili/cache"
	"wiliwili/conf"
	"wiliwili/model"
	"wiliwili/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 连接mysql
	dataMysql := conf.Data.MySql
	mysqlString := dataMysql.User + ":" + dataMysql.Password + "@tcp(" + dataMysql.Host + ":" + dataMysql.Port + ")/" +
		dataMysql.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	model.Database(mysqlString)

	// 链接redis
	dataRedis := conf.Data.Redis
	cache.Redis(dataRedis.Addr, dataRedis.Password, dataRedis.Dbname)

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
