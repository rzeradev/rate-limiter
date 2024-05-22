package ratelimiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rzeradev/rate-limiter/configs"
)

type RedisStrategy struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStrategy() *RedisStrategy {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.AppConfig.RedisAddr,
		Password: configs.AppConfig.RedisPass,
		DB:       configs.AppConfig.RedisDB,
	})

	return &RedisStrategy{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisStrategy) IsIPLimitExceeded(ip string) bool {
	key := "ip:" + ip
	return r.isLimitExceeded(key, configs.AppConfig.IPMaxReq, configs.AppConfig.BlockTime)
}

func (r *RedisStrategy) IsTokenLimitExceeded(token string) bool {
	key := "token:" + token
	return r.isLimitExceeded(key, configs.AppConfig.TokenMaxReq, configs.AppConfig.BlockTime)
}

func (r *RedisStrategy) isLimitExceeded(key string, maxReq, blockTime int) bool {
	count, err := r.client.Get(r.ctx, key).Int()
	if err != nil && err != redis.Nil {
		return true
	}

	if count >= maxReq {
		r.client.Set(r.ctx, key, count, time.Duration(blockTime)*time.Second)
		return true
	}

	r.client.Incr(r.ctx, key)
	expiration := time.Second
	if configs.AppConfig.RateLimitDur == "minute" {
		expiration = time.Minute
	}
	r.client.Expire(r.ctx, key, expiration)
	return false
}
