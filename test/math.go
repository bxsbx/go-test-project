package main

import (
	"fmt"
	"math"
	"strings"
)

func Decimals(x float64, n int) float64 {
	pow10 := math.Pow10(n)
	return math.Round(x*pow10) / pow10
}

func Tw(f float64) float64 {
	x := math.Pow10(2)
	x = math.Trunc(f*x) / x
	return x
}

func FloatToPercentage(f float64) float64 {
	f *= 100
	x := math.Pow10(2)
	x = math.Trunc(f*x) / x
	return x
}
func main() {
	for i := 1; i < 1000; i++ {
		for j := 1; j <= i; j++ {
			//decimals := Decimals(float64(j)/float64(i), 2) * 100
			decimals := FloatToPercentage(float64(j) / float64(i))
			h := Tw(decimals)
			s := fmt.Sprintf("%v", h)
			fmt.Println("---", s)
			split := strings.Split(s, ".")
			if len(split) > 1 {
				r := split[1]
				if len(r) > 2 {
					fmt.Println(s)
				}
			}
		}
	}
}
