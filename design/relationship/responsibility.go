package relationship

import "fmt"

type Handle interface {
	SetNext(handle Handle)
	Handle()
	Next()
}

type AHandle struct {
	handle Handle
	Name   string
}

func (a *AHandle) SetNext(handle Handle) {
	a.handle = handle
}

func (a *AHandle) Handle() {
	fmt.Println("执行A处理器开始")
	a.Next()
	fmt.Println("执行A处理器结束")
}

func (a *AHandle) Next() {
	if a.handle != nil {
		a.handle.Handle()
	}
}

type BHandle struct {
	handle Handle
	Name   string
	Obj    *AStruct
}

func (b *BHandle) SetNext(handle Handle) {
	b.handle = handle
}

func (b *BHandle) Handle() {
	if b.Obj.B == "B" {
		fmt.Println("执行B处理器开始")
		b.Next()
		fmt.Println("执行B处理器结束")
	} else {
		fmt.Println("无法执行B处理器")
	}
}

func (b *BHandle) Next() {
	if b.handle != nil {
		b.handle.Handle()
	}
}

type CHandle struct {
	handle Handle
	Name   string
	Obj    *AStruct
}

func (c *CHandle) SetNext(handle Handle) {
	c.handle = handle
}

func (c *CHandle) Handle() {
	if c.Obj.B == "C" {
		fmt.Println("执行C处理器开始")
		c.Next()
		fmt.Println("执行C处理器结束")
	} else {
		fmt.Println("无法执行C处理器")
	}
}

func (c *CHandle) Next() {
	if c.handle != nil {
		c.handle.Handle()
	}
}
