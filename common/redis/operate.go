package redis

import (
	"github.com/gomodule/redigo/redis"
)

// lua脚本保证锁原子性
const SetLua = `
if redis.call('set', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`

const GetLua = `
if redis.call('get', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`

// 设置一个有过期时间的键值对
func (m *redisObj) SetTTL(key string, value string, ttl int64) error {
	conn := m.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("setex", key, ttl, value); err != nil {
		return err
	}
	return nil
}

// 设置一个键值对
func (m *redisObj) Set(key string, value string) error {
	conn := m.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("set", key, value); err != nil {
		return err
	}
	return nil
}

// 获取一个键的值
func (m *redisObj) GetString(key string) (string, error) {
	conn := m.Pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}
	return value, nil
}

// 删除一个键
func (m *redisObj) Remove(key string) error {
	conn := m.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("del", key); err != nil {
		return err
	}
	return nil
}

// 判断一个键是否存在
func (m *redisObj) Exists(key string) (bool, error) {
	conn := m.Pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}
