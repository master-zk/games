package main

import (
	"games/app"
	"games/app/bundles/steam_game"
	"games/app/bundles/user"
	"github.com/gin-gonic/gin"
)

func customRegister(r *gin.Engine) {
	r.GET("/ping", app.HandlePing)
	//
	api := r.Group("/api")

	steam_game.GeneratedRegister(api)
	user.GeneratedRegister(api)
}
