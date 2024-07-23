package global

import (
	"games/app/config"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type DbMap struct {
	Jenny *gorm.DB
	Xh    *gorm.DB
}

var (
	BasePath  string              // 全局路径
	Config    *config.Config      // 配置文件
	DB        DbMap               // orm
	Validator *validator.Validate // orm
)
