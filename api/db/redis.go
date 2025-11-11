package db

import (
	"context"
	"log"
	"time"
	"workerbee/config"

	"github.com/redis/go-redis/v9"
)

func RedisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.RedisAddr,
		Password:     config.RedisPassword,
		DB:           config.RedisDB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MaxRetries:   3,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Println("Warning: Unable to connect to Redis: ", err)
		return nil
	}

	log.Println("Connected to Redis")
	return rdb
}

func IsRedisAvailable(rdb *redis.Client) bool {
	if rdb == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Println("Warning: Redis ping failed: ", err)
		return false
	}
	return true
}
