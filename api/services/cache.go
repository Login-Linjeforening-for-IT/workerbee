package services

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"workerbee/db"
	"workerbee/internal"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	redis *redis.Client
}

func NewCacheService(rdb *redis.Client) *CacheService {
	return &CacheService{redis: rdb}
}

func (cs *CacheService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if !db.IsRedisAvailable(cs.redis) {
		return nil
	}
	return cs.redis.Set(ctx, key, value, ttl).Err()
}

func (cs *CacheService) GetJSON(ctx context.Context, key string, dest any) error {
	if !db.IsRedisAvailable(cs.redis) {
		return internal.ErrCacheUnavailable
	}

	data, err := cs.redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func (cs *CacheService) SetJSONAsync(key string, value any, ttl time.Duration) {
	if !db.IsRedisAvailable(cs.redis) {
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		data, err := json.Marshal(value)
		if err != nil {
			log.Printf("Cache marshal failed for key %s: %v", key, err)
			return
		}

		if err := cs.redis.Set(ctx, key, data, ttl).Err(); err != nil {
			log.Printf("Cache set failed for key %s: %v", key, err)
		}
	}()
}

func (cs *CacheService) Delete(ctx context.Context, key string) error {
	if !db.IsRedisAvailable(cs.redis) {
		return nil
	}
	return cs.redis.Del(ctx, key).Err()
}

func (cs *CacheService) DeletePattern(ctx context.Context, pattern string) error {
	if !db.IsRedisAvailable(cs.redis) {
		return nil
	}

	iter := cs.redis.Scan(ctx, 0, pattern, 100).Iterator()
	for iter.Next(ctx) {
		cs.redis.Del(ctx, iter.Val())
	}
	return iter.Err()
}

func (cs *CacheService) Expire(ctx context.Context, key string, ttl time.Duration) error {
	if !db.IsRedisAvailable(cs.redis) {
		return nil
	}
	return cs.redis.Expire(ctx, key, ttl).Err()
}
