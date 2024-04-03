package structure

import "fmt"

//---------------静态代理----------------

type Proxy struct {
	Obj SObject
}

func (p *Proxy) AFunc() int {
	fmt.Println("AFunc pre")
	var a int
	if p.Obj != nil {
		a = p.Obj.AFunc()
		fmt.Println(a)
	}
	fmt.Println("AFunc after")
	return a
}

func (p *Proxy) BFunc() string {
	fmt.Println("BFunc pre")
	var b string
	if p.Obj != nil {
		b = p.Obj.BFunc()
		fmt.Println(b)
	}
	fmt.Println("BFunc after")
	return b
}
