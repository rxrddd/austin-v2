package redisHelper

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"math"
	"time"
)

// GetRedisCacheKey 根据参数获取redis key
func GetRedisCacheKeyByParams(module string, params map[string]interface{}) string {
	cacheKey := module + ":Administrator:"
	if len(params) == 0 {
		return cacheKey
	}
	jsonParams, _ := json.Marshal(params)
	paramsStr := string(jsonParams)
	cacheKey += paramsStr
	return cacheKey
}

// SaveRedisCache 缓存信息
func SaveRedisCache(client redis.Cmdable, key, value string) error {
	return client.Set(context.Background(), key, value, 0).Err()
}

// GetRedisCache 获取信息缓存
func GetRedisCache(client redis.Cmdable, key string) string {
	return client.Get(context.Background(), key).Val()
}

// DeleteRedisCache 删除信息缓存
func DeleteRedisCache(client redis.Cmdable, key string) {
	client.Del(context.Background(), key)
}

// BatchDeleteRedisCache 批量删除信息缓存
func BatchDeleteRedisCache(client redis.Cmdable, prefix string) {
	keys, _ := client.Scan(context.Background(), 0, prefix+"*", math.MaxInt64).Val()
	client.Del(context.Background(), keys...)
}

func MGet(ctx context.Context, rds redis.Cmdable, keys []string) (result map[string]string, err error) {
	result = make(map[string]string, 0)
	val, err := rds.MGet(ctx, keys...).Result()
	if err != nil {
		return result, nil
	}
	for i, key := range keys {
		result[key] = cast.ToString(val[i])
	}
	return result, err
}
func PipelineSetEx(ctx context.Context, rds redis.Cmdable, keys map[string]string, seconds int64) (err error) {
	_, err = rds.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		err = pipeliner.MSet(ctx, keys).Err()
		if err != nil {
			return err
		}
		for key := range keys {
			err = pipeliner.Expire(ctx, key, time.Duration(seconds)*time.Second).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
