package util

import "sync"

// 遍历list集合执行协程
func ExecuteCoroutineList(len int, f func(x int) error) error {
	var wg sync.WaitGroup
	wg.Add(len)
	anyError := make(chan error, len)
	for i := 0; i < len; i++ {
		go func(x int) {
			defer wg.Done()
			if err := f(x); err != nil {
				anyError <- err
			}
		}(i)
	}
	wg.Wait()
	close(anyError)
	for err := range anyError {
		if err != nil {
			return err
		}
	}
	return nil
}

// 遍历map执行协程
func ExecuteCoroutineMap[K comparable](AMap map[K]interface{}, f func(k K) error) error {
	var wg sync.WaitGroup
	anyError := make(chan error, len(AMap))
	for key := range AMap {
		wg.Add(1)
		go func(k K) {
			defer wg.Done()
			if err := f(k); err != nil {
				anyError <- err
			}
		}(key)
	}
	wg.Wait()
	close(anyError)
	for err := range anyError {
		if err != nil {
			return err
		}
	}
	return nil
}

type moreCoroutine struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	anyError error
}

func NewCoroutines() *moreCoroutine {
	return &moreCoroutine{}
}

func (c *moreCoroutine) Add(f func() error) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		err := f()
		if err != nil {
			c.mu.Lock()
			c.anyError = err
			c.mu.Unlock()
		}
	}()
}

func (c *moreCoroutine) Wait() error {
	c.wg.Wait()
	return c.anyError
}
