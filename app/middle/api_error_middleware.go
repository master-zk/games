package middle

import (
	"games/app/common"
	"games/app/common/api_error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiError() gin.HandlerFunc {
	return gin.CustomRecovery(recoveryHandler)
}

func recoveryHandler(ctx *gin.Context, err any) {
	if obj, ok := err.(*api_error.ApiError); ok {
		ctx.AbortWithStatusJSON(http.StatusOK, common.Response{Code: obj.Status, Message: obj.Msg})
	} else {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
