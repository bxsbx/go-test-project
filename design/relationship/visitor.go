package relationship

import "fmt"

type Element interface {
	GetName() string
	Accept(visitor Visitor)
}

type ElementA struct {
	A int
	B string
}

func (e *ElementA) GetName() string {
	return e.B
}

func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitorElementA(e)
}

type ElementB struct {
	A string
	B string
}

func (e *ElementB) GetName() string {
	return e.A
}

func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitorElementB(e)
}

type Visitor interface {
	VisitorElementA(e *ElementA)
	VisitorElementB(e *ElementB)
}

type Visitor1 struct {
	A string
}

func (v Visitor1) VisitorElementA(e *ElementA) {
	fmt.Println("1访问A元素", e.A, e.B)
}

func (v Visitor1) VisitorElementB(e *ElementB) {
	fmt.Println("1访问B元素", e.A, e.B)
}

type Visitor2 struct{}

func (v Visitor2) VisitorElementA(e *ElementA) {
	fmt.Println("2访问A元素", e.A, e.B)
}

func (v Visitor2) VisitorElementB(e *ElementB) {
	fmt.Println("2访问B元素", e.A, e.B)
}

type ObjectStructure struct {
	ElementMap map[string]Element
}

func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{ElementMap: make(map[string]Element)}
}

func (o *ObjectStructure) AddElement(e Element) {
	o.ElementMap[e.GetName()] = e
}

func (o *ObjectStructure) GetElement(key string) Element {
	return o.ElementMap[key]
}

func (o *ObjectStructure) DeleteElement(key string) {
	delete(o.ElementMap, key)
}

func (o *ObjectStructure) AcceptElement(visitor Visitor) {
	for _, v := range o.ElementMap {
		v.Accept(visitor)
	}
}
