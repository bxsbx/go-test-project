package main

import (
	"StandardProject/common/util"
	"fmt"
	"reflect"
)

type name struct {
	momo []string
}

type Mo struct {
	Id   int
	Name name
	UU   []string
}

func as(a []string) {
	fmt.Printf("---%p\n", a)
	a[0] = "33333"
	//fmt.Println(a)
}

func main() {

	//mo := Mo{
	//	23,
	//	"fwefe",
	//	[]string{"csasc", "csscs"},
	//}
	//mm := mo

	ii := []string{"csasc", "csscs"}

	aa := ii[0:1]
	aa[0] = "vfwefwe"

	fmt.Printf("%p\n", ii)
	fmt.Printf("%p\n", aa)

	fmt.Println(ii)
	fmt.Println(aa)
	as(aa)
	fmt.Println(aa)

	pp := make([]Mo, 5)

	pp[0].Id = 2323
	pp[0].Name = name{
		[]string{"2323666"},
	}
	fmt.Println(len(pp), pp)

	fmt.Printf("%p\n", &pp[0])
	fmt.Println(pp[0])
	nn := Mo{
		Id: 2345567899,
		UU: ii,
	}
	nn = pp[0]
	fmt.Printf("%p\n", nn.Name.momo)
	pp[0] = nn
	nn.Name.momo[0] = "2323111"
	fmt.Println(pp[0])
	fmt.Printf("%p\n", pp[0].Name.momo)

	dst := reflect.New(reflect.TypeOf(nn)).Elem()
	fmt.Println(dst.NumField())
	field := dst.Field(0)
	field.Set(reflect.ValueOf(23777))
	fmt.Println(dst)

	util.ColorPrint(dst, util.BLUE)

}
