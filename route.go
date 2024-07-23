package main

import (
	"games/app/bundles/game"
	"games/app/bundles/system"
	"games/app/point"
	"github.com/gin-gonic/gin"
)

func customRegister(r *gin.Engine) {
	r.GET("/ping", point.HandlePing)
	//
	api := r.Group("/api")

	system.GeneratedRegister(api)
	game.GeneratedRegister(api)
}
