package service

import (
	"context"
	"games/app/common/api_error"
	"games/app/global"
	"games/app/internal/api/request"
	"games/app/internal/model"
	"games/app/internal/repository"
	"gorm.io/gen"
)

type IUserService interface {
	CreateUser(ctx context.Context, dto request.UserCreateDTO) error
	UpdateUser(ctx context.Context, dto request.UserUpdateDTO) error
}

type UserService struct{}

// CreateUser 用户注册
func (us UserService) CreateUser(ctx context.Context, dto request.UserCreateDTO) error {
	var err error

	// 指定数据库
	repo := repository.Use(global.DB).Users
	// 检查email是否已存在
	var originUser *model.Users
	condition := []gen.Condition{
		repo.Email.Eq(dto.Email),
	}
	originUser, err = repo.WithContext(ctx).Where(condition...).First()
	if originUser != nil && originUser.ID > 0 {
		api_error.PanicApiError("email is exists")
	}
	condition1 := []gen.Condition{
		repo.Nickname.Eq(dto.Nickname),
	}
	originUser, err = repo.WithContext(ctx).Where(condition1...).First()
	if originUser != nil && originUser.ID > 0 {
		api_error.PanicApiError("nickname is exists")
	}

	// 插入数据
	err = repo.WithContext(ctx).Create(&model.Users{
		Nickname: dto.Nickname,
		Email:    dto.Email,
		Password: dto.Password,
		Status:   2, // 待激活
	})
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser 修改用户信息
func (us UserService) UpdateUser(ctx context.Context, dto request.UserUpdateDTO) error {
	var err error
	var originUser *model.Users

	repo := repository.Use(global.DB).Users
	condition := []gen.Condition{
		repo.ID.Eq(dto.Id),
	}
	originUser, err = repo.WithContext(ctx).Where(condition...).First()
	if originUser == nil {
		api_error.PanicApiError("user not exists")
	}

	condition1 := []gen.Condition{
		repo.Nickname.Eq(dto.Nickname),
	}
	originUser, err = repo.WithContext(ctx).Where(condition1...).First()
	if originUser != nil && originUser.ID != dto.Id {
		api_error.PanicApiError("nickname is exists")
	}
	updateUser := model.Users{
		Nickname: dto.Nickname,
		Avatar:   dto.Avatar,
	}
	// 修改数据
	_, err = repo.WithContext(ctx).Where(condition...).Updates(updateUser)
	if err != nil {
		api_error.PanicApiError("user update error")
	}

	return err
}
