package structure

//---------适配器模式------------

//--------------------类的适配器（没有类的概念）----------------

//--------------------对象适配器（结构组合）（常用）-------------

type Adapter interface {
	AdapterFuncA(obj SObject) int
	AdapterFuncB() int
	AdapterFuncC() int
}

type AdapterA struct {
	Obj SObject
}

// --------------通过传参方式-------------
func (a *AdapterA) AdapterFuncA(obj SObject) int {
	return obj.AFunc() * 1000
}

// --------------通过对象组合方式-----------
func (a *AdapterA) AdapterFuncB() int {
	return a.Obj.AFunc() * 1000
}

func (a *AdapterA) AdapterFuncC() int {
	return a.Obj.AFunc() + 1000
}

type AdapterB struct {
	Obj SObject
}

// --------------通过传参方式-------------
func (b *AdapterB) AdapterFuncA(obj SObject) int {
	return obj.AFunc() * 10
}

// --------------通过对象组合方式-----------
func (b AdapterB) AdapterFuncB() int {
	return b.Obj.AFunc() * 10
}

func (b AdapterB) AdapterFuncC() int {
	return b.Obj.AFunc() + 10
}

//------------------接口适配器（只实现适配接口的部分方法，go不适用）-------------------
