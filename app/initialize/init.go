package initialize

import (
	"games/app/global"
)

func AppInit() {

	OnConfigChangeInit()
}

func OnConfigChangeInit() {

	global.DB = RegisterDb(global.Config.Mysql["default"])

	global.Redis = RegisterRedis(global.Config.Redis["default"])
}
