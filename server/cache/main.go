package cache

import (
	"github.com/go-redis/redis"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis(addr, password, dbname string) {
	db, _ := strconv.ParseUint(dbname, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisClient = client
}
