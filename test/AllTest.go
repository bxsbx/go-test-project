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
type rw[T int8 | int16] interface {
	writer(t T)
	reader() T
}

type Number struct {
	buf int8
}

func (n *Number) writer(b int8) {
	n.buf = b
}

func (n *Number) reader() int8 {
	return n.buf
}

func main() {
	sum := 0
	for _, v := range "databaseName" {
		sum += int(v)
	}
	fmt.Println(sum)
}
