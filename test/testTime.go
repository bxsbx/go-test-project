package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := time.Now().Add(3 * time.Millisecond)
	fmt.Println("cascsv", end.Sub(start))
}
