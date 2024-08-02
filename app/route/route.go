package route

import (
	"games/app/internal/api/controller"
	"games/app/middle"
	"github.com/gin-gonic/gin"
)

var (
	iwt        = middle.Jwt()
	permission = middle.PermissionAuth()
)

func Register(group *gin.RouterGroup) {
	user := group.Group("/user")
	{
		userController := controller.NewUserController()
		user.GET("", userController.Info)
		user.POST("", userController.Create)
		user.PUT("", userController.Update)
		user.POST("/login", userController.Login)
		user.POST("/logout", userController.Logout)
	}
}
