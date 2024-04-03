package structure

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	adapterA1 := AdapterA{Obj: &AStruct{A: 10, B: "cas"}}
	adapterA2 := AdapterA{Obj: &BStruct{A: 100, B: "cas"}}
	a := adapterA1.AdapterFuncB()
	b := adapterA2.AdapterFuncB()
	fmt.Println(a, b)
}

func TestDecorator(t *testing.T) {
	Obj := &AStruct{A: 10, B: "AStruct"}
	decorator := &Decorator{Obj}
	decoratorA := &ConcreteDecoratorA{decorator}
	decoratorB := &ConcreteDecoratorB{decoratorA}
	AFunc := decoratorB.AFunc()
	fmt.Println(AFunc)
	bFunc := decoratorB.BFunc()
	fmt.Println(bFunc)
}

func TestProxy(t *testing.T) {
	proxy := Proxy{Obj: &AStruct{A: 10, B: "AStruct"}}
	proxy.AFunc()
	proxy.BFunc()
}

func TestFacade(t *testing.T) {
	a := &AObject{A: 10, B: "A"}
	b := &BObject{A: 100, B: "B-"}
	facade := Facade{a, b}
	add := facade.Add()
	fmt.Println(add)
	s := facade.Fmt()
	fmt.Println(s)
}

func TestBridge(t *testing.T) {
	objB := &AbstractBridgeB{B: 120}
	bridgeA := BridgeA{AInter: objB}
	fmt.Println(bridgeA.Add())
	bridgeB := BridgeB{AInter: objB}
	fmt.Println(bridgeB.Add())
}

func TestObjectTree(t *testing.T) {
	root := &Composite{Name: "root"}
	composite1 := &Composite{Name: "composite1"}
	composite2 := &Composite{Name: "composite2"}
	composite3 := &Composite{Name: "composite3"}

	leaf1 := &Leaf{Name: "leaf1"}
	leaf2 := &Leaf{Name: "leaf2"}
	leaf3 := &Leaf{Name: "leaf3"}

	composite1.Add(leaf1)
	composite1.Add(leaf2)
	composite2.Add(leaf3)

	root.Add(composite1)
	root.Add(composite2)
	root.Add(composite3)

	//fmt.Println(root.Operation())

	fmt.Println(root.GetChild(1).Operation())

}

func TestFlyweightFactory(t *testing.T) {
	factory := NewFlyweightFactory()
	sObject1 := factory.GetFactory("A").(*AStruct)
	fmt.Printf("%p\n", sObject1)
	sObject2 := factory.GetFactory("A").(*AStruct)
	fmt.Printf("%p\n", sObject2)
}
