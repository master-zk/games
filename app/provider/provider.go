package provider

import (
	"games/app/global"
	"github.com/gogf/gf/v2/os/gfile"
)

func Register() {

	// 设置全局路径
	global.BasePath = gfile.Pwd()
	// 加载配置文件
	ConfigProvider.Register()

	// mysql
	DatabaseProvider.Register()

	// 验证器
	ValidateProvider.Register()

	// redis
	RedisProvider.Register()

	// 日志
}
