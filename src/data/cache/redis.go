package cache

import (
	"context"
	"fmt"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) {
	var logger = logging.NewLogger(cfg)

	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
			Password: ""})
	ctx := context.Background()
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		logger.Fatal(logging.Redis, logging.ExternalService, err.Error(), nil)

	}
	logger.Info(logging.Redis, logging.Startup, "Connected Successfuly.", nil)
}

func GetRedsi() *redis.Client {
	return RedisClient
}
