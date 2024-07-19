package user

import (
	"games/app/bundles/user/handle/login"
	"github.com/gin-gonic/gin"
)

func GeneratedRegister(g *gin.RouterGroup) {
	v1 := g.Group("/user")
	{
		// Driver API Callback
		v1.POST("/login", login.Handle)
		v1.GET("/show", login.Handle)
	}
}
