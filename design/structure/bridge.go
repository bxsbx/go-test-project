package structure

type AbstractBridge interface {
	FuncA() int
}

type AbstractBridgeA struct {
	A int
}

func (a *AbstractBridgeA) FuncA() int {
	return a.A * 10
}

type AbstractBridgeB struct {
	B int
}

func (b *AbstractBridgeB) FuncA() int {
	return b.B + 10
}

type ObjBridge interface {
	Add() int
}

type BridgeA struct {
	AInter AbstractBridge
}

func (o *BridgeA) Add() int {
	return o.AInter.FuncA() + 100
}

type BridgeB struct {
	AInter AbstractBridge
}

func (o *BridgeB) Add() int {
	return o.AInter.FuncA() * 100
}
