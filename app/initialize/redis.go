package initialize

import (
	"games/app/config"
	"github.com/go-redis/redis"
)

func RegisterRedis(r config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		DB:       r.Database,
		Password: r.Auth,
	})
	ping := client.Ping()
	pingErr := ping.Err()
	if pingErr != nil {
		panic(pingErr)
	}

	return client
}
