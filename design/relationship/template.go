package relationship

import "fmt"

type Template struct {
	Obj SObject
}

func (t *Template) TemplateMethod1() {
	fmt.Println("set1")
	a := t.Obj.AFunc(10, 3)
	fmt.Println("set2")
	b := t.Obj.BFunc()
	fmt.Println("set3")
	fmt.Println(fmt.Sprintf("%d-%s", a, b))
}
