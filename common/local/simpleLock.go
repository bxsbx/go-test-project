package local

import "sync"

type funMap struct {
	sync.Mutex
	keyMap map[string]bool
}

type globalFunMap struct {
	sync.Mutex
	globalMap map[string]*funMap
}

var GolaLL = NewGlobalLockMap()

func NewGlobalLockMap() *globalFunMap {
	return &globalFunMap{
		globalMap: make(map[string]*funMap),
	}
}

func (g *globalFunMap) GetFunMap(key string) *funMap {
	g.Lock()
	defer g.Unlock()
	if temMap, ok := g.globalMap[key]; ok {
		return temMap
	} else {
		newFunMap := &funMap{
			keyMap: make(map[string]bool),
		}
		g.globalMap[key] = newFunMap
		return newFunMap
	}
}

func (f *funMap) SetKey(key string) bool {
	f.Lock()
	defer f.Unlock()
	if _, ok := f.keyMap[key]; ok {
		return false
	} else {
		f.keyMap[key] = true
		return true
	}
}

func (f *funMap) DelKey(key string) {
	f.Lock()
	defer f.Unlock()
	delete(f.keyMap, key)
}
