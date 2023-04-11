package cacheHepler

import (
	"austin-v2/utils/jsonHelper"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache struct {
	rds redis.Cmdable
	opt Options
}

type Options struct {
	defaultExpiration time.Duration
	err               error
}
type Option func(o *Options)

func WithErr(err error) func(ot *Options) {
	return func(ot *Options) {
		ot.err = err
	}
}
func WithExpiration(expiration time.Duration) func(ot *Options) {
	return func(ot *Options) {
		ot.defaultExpiration = expiration
	}
}

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
	err := query(ctx, v)
	if err != nil && errors.Is(err, c.opt.err) {
		c.rds.SetEX(context.Background(), key, jsonHelper.MustToString(v), time.Second*5)
		return nil
	}
	if err != nil {
		return err
	}
	c.rds.SetEX(context.Background(), key, jsonHelper.MustToString(v), c.opt.defaultExpiration)
	return nil
}
