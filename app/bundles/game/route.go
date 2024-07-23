package game

import (
	"games/app/bundles/game/game"
	"github.com/gin-gonic/gin"
)

func GeneratedRegister(g *gin.RouterGroup) {
	gameGroup := g.Group("/game")
	{
		gameGroup.POST("/query", game.Query)
		gameGroup.GET("/show", game.Show)
	}
}
