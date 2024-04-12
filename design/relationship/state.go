package relationship

import "fmt"

type State interface {
	A()
	B()
	C()
}

type AState struct {
	Val     int
	Context *Context
}

func (a *AState) A() {
	fmt.Println("执行A状态", a.Val)
	a.Context.CurState = a.Context.BState
}

func (a *AState) B() {
	fmt.Println("A状态等待执行B")
}

func (a *AState) C() {
	fmt.Println("A状态无法法执行C")
}

type BState struct {
	Val     int
	Context *Context
}

func (b *BState) A() {
	fmt.Println("B状态无法执行A")
}

func (b *BState) B() {
	fmt.Println("执行B状态", b.Val)
	b.Context.CurState = b.Context.CState
}

func (b *BState) C() {
	fmt.Println("B状态待执行C")
}

type CState struct {
	Val     int
	Context *Context
}

func (c *CState) A() {
	fmt.Println("C状态待执行A")
}

func (c *CState) B() {
	fmt.Println("C状态无法执行B")
}

func (c *CState) C() {
	fmt.Println("执行C状态", c.Val)
	c.Context.CurState = c.Context.AState
}

type Context struct {
	AState State
	BState State
	CState State

	CurState State
}

func (c *Context) A() {
	c.CurState.A()
}

func (c *Context) B() {
	c.CurState.B()
}

func (c *Context) C() {
	c.CurState.C()
}
