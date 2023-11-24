package util

import (
	"sync"
	"time"
)

// 建议使用懒加载模式
type tokenBucket struct {
	sync.Mutex
	rate              time.Duration //多久生产一个token
	maxTokenCount     int
	currentTokenCount int
	stopFlag          bool
	stopCount         int
}

func NewTokenBucket(maxTokenCount int, rate time.Duration) *tokenBucket {
	tokenBucket := &tokenBucket{
		rate:              rate,
		maxTokenCount:     maxTokenCount,
		currentTokenCount: maxTokenCount / 3,
		stopFlag:          true,
	}
	go tokenBucket.produceToken()
	return tokenBucket
}

// 单个生产
func (t *tokenBucket) produceToken() {
	for !t.stopFlag {
		t.Lock()
		if t.currentTokenCount < t.maxTokenCount {
			t.currentTokenCount++
		} else {
			t.stopCount++
			if t.stopCount > 10000 {
				t.stopFlag = true
			}
		}
		t.Unlock()
		time.Sleep(t.rate)
	}
}

// 存在多个同时消费
func (t *tokenBucket) ConsumeToken() bool {
	t.Lock()
	defer t.Unlock()
	if t.stopFlag {
		t.stopFlag = false
		go t.produceToken()
	}
	if t.currentTokenCount > 0 {
		t.currentTokenCount--
		return true
	}
	return false
}

// 可开启多个
func (t *tokenBucket) StartProduce() {
	t.Lock()
	defer t.Unlock()
	if t.stopFlag {
		t.stopFlag = false
	}
	go t.produceToken()
}

func (t *tokenBucket) StopProduce() {
	t.stopFlag = true
}
