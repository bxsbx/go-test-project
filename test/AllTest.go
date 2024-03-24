package main

import "fmt"

type Ok struct {
	A *int
}

func main() {
	ok := Ok{}
	fmt.Println(*ok.A)
}
