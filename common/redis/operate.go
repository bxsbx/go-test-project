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

//----------------------string操作------------------------------

// 设置一个有过期时间的键值对
func (m *redisObj) SetTTL(key string, value string, ttl int64) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	if _, err := conn.Do("setex", key, ttl, value); err != nil {
		return err
	}
	return nil
}

// 设置一个键值对
func (m *redisObj) Set(key string, value string) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	if _, err := conn.Do("set", key, value); err != nil {
		return err
	}
	return nil
}

// 获取一个键的值
func (m *redisObj) GetString(key string) (string, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return "", err
	}
	return redis.String(conn.Do("get", key))
}

// 删除一个键
func (m *redisObj) DelKey(key string) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	if _, err := conn.Do("del", key); err != nil {
		return err
	}
	return nil
}

// 删除多个键
func (m *redisObj) DelKeys(keys []string) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	if _, err := conn.Do("del", args...); err != nil {
		return err
	}
	return nil
}

// 判断一个键是否存在
func (m *redisObj) Exists(key string) (bool, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return false, err
	}
	return redis.Bool(conn.Do("exists", key))
}

// 获取匹配的所有键
func (m *redisObj) GetKeys(patter string) ([]string, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	return redis.Strings(conn.Do("keys", patter))
}

// 删除匹配的所有键
func (m *redisObj) DelKeysWithPatter(patter string) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	keys, err := redis.Strings(conn.Do("keys", patter))
	if err != nil {
		return err
	}
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	_, err = conn.Do("del", args...)
	return nil
}

//--------------------------list操作----------------------------

// 从list右边移除元素
func (m *redisObj) RPop(key string) (string, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return "", err
	}
	return redis.String(conn.Do("RPop", key))
}

// 从list左边移除元素
func (m *redisObj) LPop(key string) (string, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return "", err
	}
	return redis.String(conn.Do("LPop", key))
}

// 向list右边添加元素
func (m *redisObj) RPush(key string, value string) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("RPush", key, value)
	return err
}

// 向list左边添加元素
func (m *redisObj) LPush(key string, value string) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("LPush", key, value)
	return err
}

// 从list中的索引下标获取元素
func (m *redisObj) LIndex(key string, index int) (string, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return "", err
	}
	return redis.String(conn.Do("LIndex", key, index))
}

// 获取list长度
func (m *redisObj) LLen(key string) (int, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return 0, err
	}
	return redis.Int(conn.Do("LLen", key))
}

func (m *redisObj) RPopInt(key string) (int, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return 0, err
	}
	return redis.Int(conn.Do("RPop", key))
}

func (m *redisObj) RPushInt(key string, value int) error {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("RPush", key, value)
	return err
}

func (m *redisObj) LIndexInt(key string, index int) (int, error) {
	conn := m.pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return 0, err
	}
	return redis.Int(conn.Do("LIndex", key, index))
}
