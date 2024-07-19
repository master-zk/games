package app

import (
	"games/app/common/response/json"
	"github.com/gin-gonic/gin"
)

func HandlePing(ctx *gin.Context) {
	json.Success(ctx, "pong")
}
