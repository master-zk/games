package system

import (
	"games/app/bundles/system/category"
	"games/app/bundles/system/user"
	"github.com/gin-gonic/gin"
)

func GeneratedRegister(g *gin.RouterGroup) {
	systemGroup := g.Group("/system")
	{
		userGroup := systemGroup.Group("/user")
		{
			userGroup.POST("/login", user.Login)
		}
		userGroupLogin := systemGroup.Group("/user")
		userGroupLogin.Use()
		{
			// Driver API Callback
			userGroupLogin.POST("/logout", user.Logout)
		}

		categoryGroup := systemGroup.Group("/category")
		{
			// Driver API Callback
			categoryGroup.GET("/menu", category.Menu)
			categoryGroup.DELETE("/delete", category.Delete)
			categoryGroup.GET("/recommend", category.Recommend)
		}
	}
}
