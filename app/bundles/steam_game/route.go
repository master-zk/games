package steam_game

import (
	"games/app/bundles/steam_game/hello"
	"github.com/gin-gonic/gin"
)

func GeneratedRegister(g *gin.RouterGroup) {
	steamGame := g.Group("/steamGame")
	{
		// Driver API Callback
		steamGame.POST("/hello", hello.Handle)
		steamGame.GET("/hello", hello.Handle)
	}
}
