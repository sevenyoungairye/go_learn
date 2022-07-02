package redisdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"top.lel.dn/main/pkg/logger"
	"top.lel.dn/main/pkg/yaml"
)

// https://github.com/go-redis/redis
var ctx = context.Background()

// RedisCacheRepo redis ops object.
type RedisCacheRepo struct {
	Cache *redis.Client
	Ctx   context.Context
}

func New() *RedisCacheRepo {
	return &RedisCacheRepo{
		Cache: getRedisClient(),
		Ctx:   ctx,
	}
}

func (r *RedisCacheRepo) Add(k, v string, exp time.Duration) error {
	return r.Cache.Set(ctx, k, v, exp).Err()
}

func (r *RedisCacheRepo) GetVal(k string) string {
	return r.Cache.Get(ctx, k).Val()
}

func (r *RedisCacheRepo) Del(k string) {
	logger.Info("删除redis key: " + k + " 删除内容: " + r.GetVal(k))
	r.Cache.Del(ctx, k)
}

func ExampleClient() {
	rdb := getRedisClient()

	err := rdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

// GetCtx 获取上下文对象 用于额外操作redis
func GetCtx() context.Context {
	return ctx
}

// GetRedisClient get a redis client.
func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     yaml.GetRedis().Redis.Addr,
		Password: "",
		DB:       yaml.GetRedis().Redis.Db,
	})
}
