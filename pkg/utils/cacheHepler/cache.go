package cacheHepler

import (
	"austin-v2/pkg/utils/jsonHelper"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache struct {
	rds redis.Cmdable
	opt Options
}

type Options struct {
	defaultExpiration time.Duration
}
type Option func(o *Options)

func NewCache(rds redis.Cmdable, opts ...Option) *Cache {
	var o = Options{
		defaultExpiration: 30 * time.Minute,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &Cache{
		rds: rds,
		opt: o,
	}
}

type queryFun func(ctx context.Context, v interface{}) error

func (c *Cache) GetOrSet(ctx context.Context, key string, v interface{}, query queryFun) error {
	return c.GetOrSetEx(ctx, key, v, query, c.opt.defaultExpiration)
}

func (c *Cache) DelCache(ctx context.Context, keys ...string) error {
	return c.rds.Del(ctx, keys...).Err()
}

func (c *Cache) GetOrSetEx(ctx context.Context, key string, v interface{}, query queryFun, expiration time.Duration) error {
	val := c.rds.Get(context.Background(), key).Val()
	if val != "" {
		return json.Unmarshal([]byte(val), &v)
	}
	if err := query(ctx, v); err != nil {
		return err
	}
	c.rds.SetEX(context.Background(), key, jsonHelper.MustToString(v), expiration)
	return nil
}
