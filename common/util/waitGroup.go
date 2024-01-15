package util

import "sync"

type wait struct{}

type WaitGroup struct {
	sync.Mutex
	wg      sync.WaitGroup
	chanNum chan wait
	err     error
}

// 0:表示不限制协程数
func NewWaitGroup(maxNum int) *WaitGroup {
	if maxNum > 0 {
		return &WaitGroup{
			chanNum: make(chan wait, maxNum),
		}
	}
	return &WaitGroup{}
}

func (w *WaitGroup) Go(f func() error) {
	w.Add()
	go func() {
		defer w.Done()
		err := f()
		if err != nil {
			w.SetError(err)
		}
	}()
}

func (w *WaitGroup) Add() {
	w.wg.Add(1)
	if w.chanNum != nil {
		w.chanNum <- wait{}
	}
}

func (w *WaitGroup) Done() {
	w.wg.Done()
	if w.chanNum != nil {
		<-w.chanNum
	}
}

func (w *WaitGroup) Wait() error {
	w.wg.Wait()
	return w.err
}

func (w *WaitGroup) SetError(err error) {
	w.Lock()
	w.err = err
	w.Unlock()
}

//func (w *WaitGroup) GetError() error {
//	return w.err
//}
