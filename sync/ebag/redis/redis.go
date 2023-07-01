package redis

import (
	"fmt"
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type redisObj struct {
	pool *redis.Pool
}

type redisConfig struct {
	Server      string
	Password    string
	DBNum       int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Wait        bool
}

var RedisObj *redisObj

const (
	ZS_QUESTIONS   = "sync::get::qs::zs::"
	EBAG_QUESTIONS = "sync::get::qs::ebag::"

	FAIL_ZS_QUESTIONS   = "sync::fail::qs::zs::"
	FAIL_EBAG_QUESTIONS = "sync::fail::qs::ebag::"

	SQL_INSERT = "sync::sql::qs::fail::insert"
	SQL_UPDATE = "sync::sql::qs::fail::update"

	STUDENT_SCHOOL = "sync::stu::school::"

	//记录执行成功的页数
	RECORD_SUCCESS_PAGE = "sync::success::page"
)

func defaultRedisConfig(cfg beegoConfig.Configer) redisConfig {
	dbNum, _ := cfg.Int("SyncConfig::RedisDBNum")
	maxIdle, _ := cfg.Int("SyncConfig::RedisMaxIdle")
	maxActive, _ := cfg.Int("SyncConfig::RedisMaxActive")
	idleTimeout, _ := cfg.Int("SyncConfig::RedisIdleTimeout")
	wait, _ := cfg.Bool("SyncConfig::RedisWait")
	return redisConfig{
		Server:      cfg.String("SyncConfig::RedisServer"),
		Password:    cfg.String("SyncConfig::RedisPassword"),
		DBNum:       dbNum,
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Wait:        wait,
	}
}

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
					fmt.Println(err)
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
	RedisObj = newRedis(defaultRedisConfig(cfg))
}
