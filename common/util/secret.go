package util

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// appId 取16位字符串，appSecret 取32位
func GenerateAppIdAndSecret(appName string) (string, string) {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(base64.URLEncoding.EncodeToString([]byte(appName)))))
	return md5Str[:16], md5Str[16:48]
}
