package util

import "sync"

type WaitGroup struct {
	mu         sync.Mutex
	wg         sync.WaitGroup
	chanNum    chan bool
	maxLineNum int
	err        error
}

func NewWaitGroup(maxLineNum int) *WaitGroup {
	return &WaitGroup{
		chanNum: make(chan bool, maxLineNum),
	}
}

func (w *WaitGroup) Add() {
	w.wg.Add(1)
	w.chanNum <- true
}

func (w *WaitGroup) Done() {
	w.wg.Done()
	<-w.chanNum
}

func (w *WaitGroup) Wait() {
	w.wg.Wait()
}

func (w *WaitGroup) Lock() {
	w.mu.Lock()
}

func (w *WaitGroup) Unlock() {
	w.mu.Unlock()
}

func (w *WaitGroup) SetError(err error) {
	w.err = err
}

func (w *WaitGroup) GetError() error {
	return w.err
}
