package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

func CloseRedis() {
	RedisClient.Close()
}

func Set[T any](c *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx := context.Background()
	return c.Set(ctx, key, v, duration).Err()
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var dest T = *new(T)
	ctx := context.Background()

	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
