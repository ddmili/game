package redis

import (
	"time"

	rv9 "github.com/redis/go-redis/v9"
)

var (
//	RedisCtx = context.Background()
)



// NewPoolRedis 获取连接池
func NewPoolRedis(addr string, password string) *rv9.Client {
	return rv9.NewClient(&rv9.Options{
		Addr:         addr,
		Password:     password, // no password set
		DB:           0,        // use default DB
		MinIdleConns: 20,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     30,
		PoolTimeout:  100 * time.Second,
		MaxRetries:   3,
	})
}