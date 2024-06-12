package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ParasHarnagle/verve_task/models"
	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "my-redis-container:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis server: %v\n", err)
		return err
	}

	log.Println("Connected to Redis server successfully")
	return nil
}

func GetPromotionFromCache(id string) (models.Promotion, error) {
	val, err := redisClient.Get(ctx, id).Bytes()
	if err == redis.Nil {
		return models.Promotion{}, fmt.Errorf("not in cache")
	} else if err != nil {
		return models.Promotion{}, fmt.Errorf("error: %v", err)
	}
	var p models.Promotion
	if err := json.Unmarshal(val, &p); err != nil {
		return models.Promotion{}, err
	}

	return p, nil
}

func PromotionToCache(p models.Promotion) error {
	val, err := json.Marshal(p)
	if err != nil {
		return err
	}
	expiration := 30 * time.Minute
	log.Printf("Promotion with ID %s cached successfully\n", p.ID)
	err = redisClient.Set(ctx, p.ID, val, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func ClearCache() error {
	err := redisClient.FlushAll(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
