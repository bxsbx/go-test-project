package util

import "strings"

// 生成重复字符串
func GenerateDuplicates(count int, src, suffix string) string {
	repeat := strings.Repeat(src, count)
	return strings.TrimSuffix(repeat, suffix)
}
