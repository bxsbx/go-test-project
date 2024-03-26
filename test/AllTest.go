package main

import (
	"StandardProject/common/cryptoes"
	"fmt"
	"unicode/utf8"
)

// 脱敏姓名
func desensitizeName(name string) string {
	runeCount := utf8.RuneCountInString(name)
	if runeCount == 0 {
		return ""
	}
	first, _ := utf8.DecodeRuneInString(name)
	desensitizedName := string(first)
	for i := 1; i < runeCount-1; i++ {
		desensitizedName += "*"
	}
	if runeCount > 2 {
		last, _ := utf8.DecodeLastRuneInString(name)
		desensitizedName += string(last)
	} else {
		desensitizedName += "*"
	}
	return desensitizedName
}

func main() {
	originalName := "440983199901081219"
	fmt.Println("Original name:", originalName)

	desensitizedName := cryptoes.DesensitizeIdentityCard(originalName)
	fmt.Println("Desensitized name:", desensitizedName)
}
