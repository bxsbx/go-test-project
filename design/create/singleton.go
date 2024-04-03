package create

import "sync"

//-------------单例模式-------------

type Single struct {
	A int
}

// 预加载

var (
	MySingle = &Single{}
	lock     sync.Mutex
)

// 懒加载
// 懒加载不浪费内存，但是无法保证线程的安全
func GetSingle() *Single {
	if MySingle == nil {
		MySingle = &Single{}
	}
	return MySingle
}

// 缺点：每次都加锁，影响性能。
func GetSafeSingle() *Single {
	lock.Lock()
	defer lock.Unlock()
	if MySingle == nil {
		MySingle = &Single{}
	}
	return MySingle
}

// 不用每次加锁，获取前加锁，获取后不用加锁（双重if加锁方式）
func GetSafeSingleGood() *Single {
	if MySingle == nil {
		lock.Lock()
		defer lock.Unlock()
		if MySingle == nil {
			MySingle = &Single{}
		}
		return MySingle
	}
	return MySingle
}
