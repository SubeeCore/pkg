package redis

import (
	"context"

	"github.com/alyxcore/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type cacheClient struct {
	rdb redis.UniversalClient
}

func NewRedisCache(ctx context.Context, rdb redis.UniversalClient) cache.Cache {
	return &cacheClient{rdb: rdb}
}
