package main

import (
	"games/app/global"
	"games/app/initialize"
	"games/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	defer recycle()
	// 加载配置文件
	config.LoadGlobalConfig()
	// 初始化服务
	initialize.AppInit()

	// 设置运行模式
	gin.SetMode(global.Config.App.Mode)
	// 创建一个gin引擎
	engine := gin.New()
	// 设置一个信任ip : 127.0.0.1
	_ = engine.SetTrustedProxies([]string{"127.0.0.1"})
	// 路由配置
	configRoute(engine)

	// 开启http服务
	_ = engine.Run(":" + global.Config.Server.Port)
}

// 失败回调
func recycle() {
	log.Println("recycling......")
}
