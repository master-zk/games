package common

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"net/http"
)

const (
	YES int = 1
	NO  int = 2
)

type ListResponse struct {
	Total       int64 `json:"total" example:"99"`
	PageSize    int64 `json:"pageSize" example:"10"`
	CurrentPage int64 `json:"currentPage" example:"1"`
	MaxPage     int64 `json:"maxPage" example:"10"`
	Data        int64 `json:"data" example:"10"`
}

type EnumIntItemResponse struct {
	Title string `json:"title" example:"int枚举1"`
	Value int64  `json:"value" example:"1"`
	Icon  string `json:"icon" example:"https://*****.png"`
}

type EnumStringItemResponse struct {
	Title string `json:"title" example:"string枚举1"`
	Value string `json:"value" example:"code_gen"`
	Icon  string `json:"icon" example:"https://*****.png"`
}

const SuccessMessage = "OK"

type Response struct {
	Code      int    `json:"code" example:"0"`
	Message   string `json:"message" example:"OK"`
	Timestamp int64  `json:"timestamp" example:"1721006243"`
	Data      any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	r := Response{
		Code:      OK,
		Message:   SuccessMessage,
		Timestamp: carbon.Now().Timestamp(),
		Data:      data,
	}

	Json(c, http.StatusOK, r)
}

func Json(c *gin.Context, code int, response Response) {
	c.JSON(code, response)
}

func Fail(c *gin.Context, message string) {
	r := Response{
		Code:      MessageError,
		Message:   message,
		Timestamp: carbon.Now().Timestamp(),
		Data:      nil,
	}

	Json(c, http.StatusOK, r)
}
