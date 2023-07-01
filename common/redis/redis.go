package redis

import (
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

const (
	DEFAULT = "default"
)

type redisObj struct {
	pool *redis.Pool
}

var redisMap = make(map[string]*redisObj)

func newRedis(cfg redisConfig) *redisObj {
	return &redisObj{
		pool: &redis.Pool{
			MaxIdle:     cfg.MaxIdle,
			MaxActive:   cfg.MaxActive,
			IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
			Wait:        cfg.Wait,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", cfg.Server)
				if err != nil {
					log.Fatalf("tcp:%v", err)
					return c, err
				}

				_, err = c.Do("AUTH", cfg.Password)
				if err != nil {
					log.Fatalf("err password:%v", cfg.Password)
					c.Close()
					return c, err
				}

				_, err = c.Do("SELECT", cfg.DBNum)
				if err != nil {
					c.Close()
					return c, err
				}
				return c, err
			},
		},
	}
}

func InitRedis(cfg beegoConfig.Configer) {
	redisMap[DEFAULT] = newRedis(defaultRedisConfig(cfg))
}

func DefaultRedisObj() *redisObj {
	return redisMap[DEFAULT]
}

func GetRedisObj(key string) *redisObj {
	return redisMap[key]
}
