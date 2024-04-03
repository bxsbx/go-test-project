package main

import (
	"fmt"
	"sync"
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

type Al struct {
	sync.RWMutex
	A int
}

// NVUZCk1HwYEovkhSQhDwonCFsIsGgRom
func main() {
	al := Al{A: 10}
	al.RLock()
	defer al.RUnlock()
	fmt.Println(al.A)

	al.Lock()
	defer al.Unlock()
	al.A = 100
	fmt.Println(al.A)
}
