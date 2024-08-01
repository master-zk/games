package provider

import (
	"games/app/global"
	"github.com/go-redis/redis/v8"
)

type redisProvider struct {
}

var RedisProvider = &redisProvider{}

func (p *redisProvider) Register() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Host + ":" + global.Config.Redis.Port,
		DB:       global.Config.Redis.Database,
		Password: global.Config.Redis.Auth,
	})
}
