package util

import "strings"

// 生成重复字符串
func GenerateDuplicates(count int, src, suffix string) string {
	repeat := strings.Repeat(src, count)
	return strings.TrimSuffix(repeat, suffix)
}

func FirstLower(s string) string {
	if len(s) > 1 {
		return strings.ToLower(s[:1]) + s[1:]
	}
	return strings.ToLower(s)
}
