package util

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	size  int
	maps  []map[K]V
	locks []sync.RWMutex
}

func NewSafeMap[K comparable, V any](n int) *SafeMap[K, V] {
	sm := &SafeMap[K, V]{}
	sm.size = n
	sm.maps = make([]map[K]V, n)
	sm.locks = make([]sync.RWMutex, n)
	for i := 0; i < n; i++ {
		sm.maps[i] = make(map[K]V)
	}
	return sm
}

func (sm *SafeMap[K, V]) ReadMap(key K) V {
	index := hash(key) % sm.size
	sm.locks[index].RLock()
	value := sm.maps[index][key]
	sm.locks[index].RUnlock()
	return value
}

func (sm *SafeMap[K, V]) WriteMap(key K, value V) {
	index := hash(key) % sm.size
	sm.locks[index].Lock()
	sm.maps[index][key] = value
	sm.locks[index].Unlock()
}

func hash[K comparable](key K) int {
	keyStr := fmt.Sprintf("%v", key)
	h := 0
	for i := 0; i < len(keyStr); i++ {
		h += int(keyStr[i])
	}
	return h
}
