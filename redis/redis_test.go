package redis_test

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

var (
	mu sync.Mutex
	c  = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   8,
	})
)

func init() {
	if err := c.Ping().Err(); err != nil {
		log.Fatalln(err)
	}
}

func safeFlushDB() {
	mu.Lock()
	defer mu.Unlock()
	c.FlushDB()
}
