package util

import (
	"crypto/md5"
	"fmt"
)

// md5
func Md5(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}
