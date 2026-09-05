package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) {
	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
			Password: ""})
	ctx := context.Background()
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)

	}
	log.Println("Redis connected successfully")
}

func GetRedsi() *redis.Client {
	return RedisClient
}
