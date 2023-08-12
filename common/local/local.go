package local

import (
	"sync"
	"time"
)

type lockMutex struct {
	mu          *sync.Mutex
	expiredTime time.Time
}

// 仅适用于单机，分布式锁可以用redis
type lockMap struct {
	cycleTime int // 秒 执行循环删除的时间
	expired   int // 秒 过期时间
	sync.Mutex
	muMap map[string]*lockMutex
}

func NewLockMap(expired, cycleTime int) *lockMap {
	lock := &lockMap{
		muMap:     make(map[string]*lockMutex),
		expired:   expired,
		cycleTime: cycleTime,
	}
	go lock.startCleanUpRoutine()
	return lock
}

func (t *lockMap) GetLock(key string) *sync.Mutex {
	t.Lock()
	defer t.Unlock()
	if lock, ok := t.muMap[key]; ok {
		return lock.mu
	}

	lock := lockMutex{
		mu:          &sync.Mutex{},
		expiredTime: time.Now().Add(time.Duration(t.expired) * time.Second),
	}
	t.muMap[key] = &lock
	return lock.mu
}

// 删除锁(避免产生的锁过多而占用内存)
func (t *lockMap) DeleteExpiredLock() {
	t.Lock()
	defer t.Unlock()
	now := time.Now()
	for k, lock := range t.muMap {
		if now.After(lock.expiredTime) {
			delete(t.muMap, k)
		}
	}
}

func (t *lockMap) startCleanUpRoutine() {
	ticker := time.NewTicker(time.Duration(t.cycleTime) * time.Second)
	for range ticker.C {
		t.DeleteExpiredLock()
	}
}
