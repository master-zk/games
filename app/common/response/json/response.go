package json

import (
	"fmt"
	"games/app/common/response/error_code"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"net/http"
	"reflect"
)

const SuccessMessage = "OK"

type Response struct {
	Code      int    `json:"code" example:"0"`
	Message   string `json:"message" example:"OK"`
	Timestamp int64  `json:"timestamp" example:"1721006243"`
	Data      any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	r := Response{
		Code:      error_code.Success,
		Message:   SuccessMessage,
		Timestamp: carbon.Now().Timestamp(),
		Data:      data,
	}
	fmt.Println(r)

	Json(c, http.StatusOK, r)
}

func Json(c *gin.Context, code int, response Response) {
	fmt.Printf("%T, %v\n", response, reflect.TypeOf(response).Kind())
	c.JSON(code, response)
}

func Fail(c *gin.Context, message string) {
	r := Response{
		Code:      error_code.MessageError,
		Message:   message,
		Timestamp: carbon.Now().Timestamp(),
		Data:      nil,
	}

	Json(c, http.StatusOK, r)
}
