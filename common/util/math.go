package util

import "math"

// 保留n位小数
func Decimals(x float64, n int) float64 {
	pow10 := math.Pow10(n)
	return math.Round(x*pow10) / pow10
}
