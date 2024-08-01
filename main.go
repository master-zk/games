package main

import (
	"games/app/global"
	"games/app/provider"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	defer recycle()

	// 注册服务
	provider.Register()

	// 注册路由
	router := gin.Default()
	customRegister(router)

	err := router.Run(":" + global.Config.Server.Port)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}

// 失败回调
func recycle() {
	log.Println("recycling......")
}
