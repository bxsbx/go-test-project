package create

import "fmt"

//---------------------具体对象-------------------

type SObject interface {
	AFunc()
	BFunc(s string)
	CFunc(n int)
}

//---------------对象A-----------------

type AStruct struct {
	A int
	B string
}

func (a *AStruct) AFunc() {
	fmt.Println("AStruct")
}

func (a *AStruct) BFunc(s string) {
	fmt.Println("AStruct", s)
}

func (a *AStruct) CFunc(n int) {
	fmt.Println("AStruct", n)
}

//---------------对象B-----------------

type BStruct struct {
	B int
	A string
}

func (b *BStruct) AFunc() {
	fmt.Println("BStruct")
}

func (b *BStruct) BFunc(s string) {
	fmt.Println("BStruct", s)
}

func (b *BStruct) CFunc(n int) {
	fmt.Println("BStruct", n)
}

type LObject interface {
	AFunc(s string)
	BFunc(n int)
}

//---------------对象C-----------------

type CStruct struct {
	A int
	B string
}

func (c *CStruct) AFunc(s string) {
	fmt.Println("CStruct", s)
}

func (c *CStruct) BFunc(n int) {
	fmt.Println("CStruct", n)
}

//---------------对象D-----------------

type DStruct struct {
	B int
	A string
}

func (d *DStruct) AFunc(s string) {
	fmt.Println("BStruct", s)
}

func (d *DStruct) BFunc(n int) {
	fmt.Println("BStruct")
}
