package middle

import (
	"github.com/gin-gonic/gin"
)

func PermissionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullPath := c.FullPath()
		if fullPath != "" {

		}
		c.Next()
	}
}
