package main

import (
	"games/app/common"
	"games/app/middle"
	"games/app/point"
	"games/app/route"
	"github.com/gin-gonic/gin"
)

func configRoute(engine *gin.Engine) {
	engine.Use(
		gin.Logger(),
		//gin.Recovery(),
		middle.Cors(),
		middle.ApiError(),
	)

	// 健康检测
	engine.GET("/ping", point.HandlePing)
	// 未注册路由
	engine.NoRoute(func(c *gin.Context) {
		common.CodeResponse(c, common.CodeNoRoute)
	})

	// api路由组
	group := engine.Group("/api")
	{
		route.Register(group)
	}
}
