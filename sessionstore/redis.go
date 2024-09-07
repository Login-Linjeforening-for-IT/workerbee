package sessionstore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string `config:"SESSION_STORE_ADDR" default:"localhost:6379"`
	Username string `config:"SESSION_STORE_USERNAME" default:""`
	Password string `config:"SESSION_STORE_PASSWORD" default:""`
	DB       int    `config:"SESSION_STORE_DB" default:"0"`
}

type redisStore struct {
	client *redis.Client
}

func New(client *redis.Client) Store {
	return &redisStore{
		client: client,
	}
}

func getRedisSessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func (store *redisStore) CreateSession(ctx context.Context, params CreateSessionParams) error {
	session := Session{
		ID:           params.ID,
		UID:          params.UID,
		RefreshToken: params.RefreshToken,
		UserAgent:    params.UserAgent,
		ClientIP:     params.ClientIP,
		ExpiresAt:    params.ExpiresAt,

		CreatedAt: time.Now(),
	}

	return store.setSession(ctx, session, time.Until(session.ExpiresAt))
}

func (store *redisStore) setSession(ctx context.Context, session Session, expiration time.Duration) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}

	id := getRedisSessionKey(session.ID)

	return store.client.Set(ctx, id, string(data), expiration).Err()
}

func (store *redisStore) GetSession(ctx context.Context, id string) (Session, error) {
	data, err := store.client.Get(ctx, fmt.Sprintf("session:%s", id)).Result()
	if err != nil {
		if err == redis.Nil {
			return Session{}, &NotFoundError{
				Key:      id,
				Resource: "sessions",
				Err:      err,
			}
		}
		return Session{}, err
	}

	var session Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return Session{}, err
	}

	return session, nil
}

func (store *redisStore) DeleteSession(ctx context.Context, id string) error {
	return store.client.Del(ctx, getRedisSessionKey(id)).Err()
}

func (store *redisStore) BlockSession(ctx context.Context, id string) error {
	res, err := store.GetSession(ctx, id)
	if err != nil {
		return err
	}
	res.BlockedAt = time.Now()

	return store.setSession(ctx, res, 0) // No expiration for future reference
}
