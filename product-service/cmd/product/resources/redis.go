package resources

import (
	"context"
	"log"
	"product-service/config"
)
import "github.com/redis/go-redis/v9"

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
	})

	ctx := context.Background()
	pingResults, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected:", pingResults)
	return RedisClient
}
