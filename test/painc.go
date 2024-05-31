package main

import "fmt"

func Slice(list []int) {
	fmt.Printf("The address of x is: %p\n", &list)
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
	fmt.Println(cap(list))
	fmt.Printf("The address of x is: %p\n", &list)

}

func main() {
	list := make([]int, 10)
	Slice(list)
	fmt.Println(cap(list))
}
