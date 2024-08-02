package global

import (
	"games/app/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	BasePath string         // 全局路径
	Config   *config.Config // 配置文件
	DB       *gorm.DB       // gorm
	Redis    *redis.Client  // redis
)
