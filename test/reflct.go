package main

import (
	"fmt"
	"reflect"
)

type One1 struct {
	Id    string
	Name  string
	Count int
}

func main() {
	var one One1
	fmt.Println(one)
	of := reflect.TypeOf(one)
	fmt.Println(of)
	fmt.Println(of.NumField())
	//fmt.Println(of.Elem().String())
	fmt.Println(of.Field(0).Name)

	value := reflect.ValueOf(one)
	fmt.Println(value)
	fmt.Println(value.NumField())
	//fmt.Println(of.Elem().String())
	fmt.Println(value.Field(2))

	//var aa []string
	//typeOf := reflect.TypeOf(aa)
	//fmt.Println(typeOf)
	//valueOf := reflect.ValueOf(aa)
	//field := valueOf.Field(0)
	//fmt.Println(field)
}
