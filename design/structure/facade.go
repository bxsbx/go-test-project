package structure

import "fmt"

type AObject struct {
	A int
	B string
}

func (a *AObject) AddA() int {
	return a.A + 10
}

func (a *AObject) FmtA() string {
	return a.B + "FmtA"
}

type BObject struct {
	A int
	B string
}

func (b *BObject) AddB() int {
	return b.A + 10
}

func (b *BObject) FmtB() string {
	return b.B + "FmtB"
}

type Facade struct {
	A *AObject
	B *BObject
}

func (f *Facade) Add() int {
	return f.A.AddA() + f.B.AddB()
}

func (f *Facade) Fmt() string {
	return fmt.Sprintf("%d-%s", f.A.AddA(), f.B.FmtB())
}
