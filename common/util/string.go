package util

import "strings"

// 生成重复字符串
func GenerateDuplicates(count int, src, suffix string) string {
	repeat := strings.Repeat(src, count)
	return strings.TrimSuffix(repeat, suffix)
}

func FirstLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

// 驼峰命名
func HumpNaming(s string) string {
	split := strings.Split(s, "_")
	for i := 0; i < len(split); i++ {
		str := split[i]
		if len(str) > 0 {
			split[i] = strings.ToUpper(str[:1]) + str[1:]
		}
	}
	return strings.Join(split, "")
}
