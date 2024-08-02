package controller

import (
	"games/app/common"
	"games/app/internal/api/request"
	"games/app/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.IUserService
}

func NewUserController() UserController {
	return UserController{service: service.UserService{}}
}

func (ctl UserController) Login(c *gin.Context) {
	common.Success(c, nil)
}

func (ctl UserController) Create(c *gin.Context) {
	var dto request.UserCreateDTO
	if err := c.ShouldBind(&dto); err != nil {
		common.CodeResponse(c, common.CodeApiParamsError)
	}

	err := ctl.service.CreateUser(c.Request.Context(), dto)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}

	common.Success(c, nil)
}

func (ctl UserController) Update(c *gin.Context) {
	var dto request.UserUpdateDTO
	if err := c.ShouldBind(&dto); err != nil {
		common.CodeResponse(c, common.CodeApiParamsError)
	}

	err := ctl.service.UpdateUser(c.Request.Context(), dto)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}

	common.Success(c, nil)
}

func (ctl UserController) Logout(c *gin.Context) {
	common.Success(c, nil)
}

func (ctl UserController) Info(c *gin.Context) {
	common.Success(c, nil)
}
