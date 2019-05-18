package redis

import (
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type Rediser interface {
	Incr(key string) *redis.IntCmd
	Decr(key string) *redis.IntCmd
	Get(key string) *redis.StringCmd
}

var once sync.Once

var TestingRDB Rediser
var rdb Rediser

var dsn string

func SetDSN(s string) {
	dsn = s
}

func GetClient() (Rediser, error) {
	if flag.Lookup("test.v") != nil && TestingRDB != nil {
		return TestingRDB, nil
	}

	if dsn == "" {
		return nil, errors.New("Redis DSN not set")
	}

	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:         ":6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
		})
	})

	return rdb, nil
}

func GetKey(s string) string {
	return fmt.Sprintf("pp:%s", s)
}
