package hello

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context) {
	fmt.Println("Hello World")
}
