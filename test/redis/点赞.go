package main

import (
	"server/core"
	"server/global"
	"server/service/redis_ser"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()
	redis_ser.Digg("xxx")
}
