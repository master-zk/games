package login

import (
	"games/app/common/response/json"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Errno     int    `json:"errno" example:"200"`
	Errmsg    string `json:"errmsg" example:"status bad request"`
	Timestamp string `json:"timestamp"`
	Data      any    `json:"data,omitempty" example:"null"`
}

func Handle(c *gin.Context) {
	json.Success(c, nil)
}
