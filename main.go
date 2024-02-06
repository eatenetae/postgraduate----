package main

import (
	"server/core"
	_ "server/docs"
	"server/flag"
	"server/global"
	"server/routers"
	"server/service/cron_ser"
)

// @title MyBlog API文档
// @version 1.0
// @description MyBlog API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 连接redis
	global.Redis = core.ConnectRedis()
	// 连接es
	global.ESClient = core.EsConnect()
	// 连接ip地址数据库
	core.InitAddrDB()
	defer global.AddrDB.Close()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// 定时任务
	cron_ser.CronInit()

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("my_blog_server运行于:  %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
