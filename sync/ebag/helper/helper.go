package helper

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var TIMELOCAL *time.Location

func init() {
	local, _ := time.LoadLocation("Asia/Chongqing") //服务器设置的时区
	TIMELOCAL = local
}

// md5
func Md5(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}

// 把一个时间字符串转为unix时间戳
func StrToTimeStamp(timeStr string) int64 {
	//	time = "2015-09-14 16:33:00"
	loc, _ := time.LoadLocation("Asia/Chongqing")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return t.Unix()
}

// get now datatime(Y-m-d H:i:s)
func GetNowDateTime() string {
	return time.Now().In(TIMELOCAL).Format("2006-01-02 15:04:05")
}

func StringToIntArray(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}

// JSON string
func JSONStringfy(data interface{}) string {
	jb, _ := json.Marshal(data)
	return string(jb)
}
