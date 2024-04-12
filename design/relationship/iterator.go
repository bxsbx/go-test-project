package relationship

type Iterator interface {
	HasNext() bool
	GetNext() AStruct
}

type AIterator struct {
	Index int
	List  []AStruct
}

func (a *AIterator) HasNext() bool {
	return a.Index < len(a.List)
}

func (a *AIterator) GetNext() AStruct {
	if a.HasNext() {
		item := a.List[a.Index]
		a.Index++
		return item
	}
	return AStruct{}
}

type BIterator struct {
	Index int
	Keys  []string
	Map   map[string]AStruct
}

func (b *BIterator) HasNext() bool {
	return b.Index < len(b.Keys)
}

func (b *BIterator) GetNext() AStruct {
	if b.HasNext() {
		item := b.Map[b.Keys[b.Index]]
		b.Index++
		return item
	}
	return AStruct{}
}

// 一般来说都是一个集合对应一个迭代器，java中的集合会常用迭代器（带泛型）

type Collection interface {
	GetIterator() Iterator
}

type ACollection struct {
	List []AStruct
}

func (a *ACollection) GetIterator() Iterator {
	return &AIterator{Index: 0, List: a.List}
}

func (a *ACollection) Add(item AStruct) {
	a.List = append(a.List, item)
}

type BCollection struct {
	Map map[string]AStruct
}

func (b *BCollection) GetIterator() Iterator {
	var keys []string
	for k := range b.Map {
		keys = append(keys, k)
	}
	return &BIterator{Index: 0, Map: b.Map, Keys: keys}
}

func (b *BCollection) Add(item AStruct) {
	b.Map[item.B] = item
}
