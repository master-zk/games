package point

import (
	"games/app/common"
	"github.com/gin-gonic/gin"
)

func HandlePing(ctx *gin.Context) {
	common.Success(ctx, "pong")
}
