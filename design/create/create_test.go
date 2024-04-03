package create

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	factory := AbstractFactory{}
	//lObject := factory.GetFactory("A").CreateLObject("C")
	//if lObject != nil {
	//	lObject.AFunc("sas")
	//}
	object := factory.CreatBFactory().CreateLObjectStruct()
	object.BFunc(12)
}

func TestBuilder(t *testing.T) {
	a := new(ABuilder)
	b := new(BBuilder)
	var d Director
	aa := d.GetProductFromBuilder(a)
	fmt.Println(aa)
	bb := d.GetProductFromBuilder(b)
	fmt.Println(bb)
}
