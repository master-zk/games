package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	customRegister(router)

	err := router.Run("127.0.0.1:8888")
	if err != nil {
		return
	}
}
