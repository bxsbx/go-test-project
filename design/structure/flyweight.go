package structure

import "sync"

type FlyweightFactory struct {
	sync.RWMutex
	HashMap map[string]SObject
}

func NewFlyweightFactory() *FlyweightFactory {
	factory := FlyweightFactory{}
	factory.HashMap = make(map[string]SObject)
	return &factory
}

func (f *FlyweightFactory) GetFactory(key string) SObject {
	if v, ok := f.HashMap[key]; ok {
		return v
	}
	f.Lock()
	defer f.Unlock()
	if v, ok := f.HashMap[key]; ok {
		return v
	}
	aObject := &AStruct{B: key}
	f.HashMap[key] = aObject
	return aObject
}
