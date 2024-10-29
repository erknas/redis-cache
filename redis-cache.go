package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const ttl = time.Second * 5

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, bool) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}

	return val, true
}

func (c *RedisCache) Set(ctx context.Context, key, val string) error {
	_, err := c.client.Set(ctx, key, val, ttl).Result()
	return err
}

func (c *RedisCache) Remove(ctx context.Context, key string) error {
	_, err := c.client.Del(ctx, key).Result()
	return err
}
