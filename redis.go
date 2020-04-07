package connect

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

type redisConfig struct {
	host     string
	port     string
	password string
	redis    *redis.Client
}

var (
	redisClient *redis.Client
	redisConf *redisConfig
)

func Redis() *redis.Client {
	if redisClient == nil {
		if redisConf.redis == nil {
			_ = redisConf.Connect()
		}
		redisClient = redisConf.redis
	}
	return redisClient
}

func NewRedisConfig(host, port, password string) *redisConfig {
	redisConf =  &redisConfig{host: host, port: port, password: password}
	return redisConf
}

func (r *redisConfig) Connect() error {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", r.host, r.port),
			Password:     r.password,
			DB:           0, // use default DB
			DialTimeout:  5 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		})
	}
	r.redis = redisClient
	fmt.Println(r.redis.Set("a", "123", 0))
	fmt.Println(r.redis.Get("a"))
	return nil
}

func (r *redisConfig) Close() error {
	return r.redis.Close()
}
