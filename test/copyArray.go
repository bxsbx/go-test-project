package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a := []int{23, 239}

	c := make(map[string]string)
	c["1"] = "2"
	c["3"] = "223"
	//b := []int{203}
	//reflect.ValueOf(&a[0]).Elem().Set(reflect.ValueOf(b[0]))
	//i := reflect.ValueOf(a[0]).Interface()
	//reflect.ValueOf(&i).Elem().Set(reflect.ValueOf(b[0]))

	//elem := reflect.ValueOf(&a).Elem()
	//i := elem.Len()
	//fmt.Println(i)
	//fmt.Println(elem.Index(0))
	//fmt.Println(a)

	//elem := reflect.ValueOf(&c).Elem()
	//i := elem.Len()
	//fmt.Println(i)
	////fmt.Println(elem.MapIndex(elem.MapKeys()[1]))
	//next := elem.MapRange().Next()
	//fmt.Println(next)
	//fmt.Println(elem.Type())
	//fmt.Println(c)

	var jj = []string{"23235"}
	slice := reflect.MakeSlice(reflect.TypeOf(jj), 0, 0)

	slice = reflect.Append(slice, reflect.ValueOf(jj[0]))
	fmt.Println(slice)

	for k := range c {
		fmt.Println(k)
	}

}
