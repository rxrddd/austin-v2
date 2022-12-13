package redisHelper

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"math"
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
	return client.Set(key, value, 0).Err()
}

// GetRedisCache 获取信息缓存
func GetRedisCache(client redis.Cmdable, key string) string {
	return client.Get(key).Val()
}

// DeleteRedisCache 删除信息缓存
func DeleteRedisCache(client redis.Cmdable, key string) {
	client.Del(key)
}

// BatchDeleteRedisCache 批量删除信息缓存
func BatchDeleteRedisCache(client redis.Cmdable, prefix string) {
	keys, _ := client.Scan(0, prefix+"*", math.MaxInt64).Val()
	client.Del(keys...)
}
