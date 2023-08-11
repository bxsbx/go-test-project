package global

import (
	"StandardProject/common/zaplog"
	"fmt"
)

func Init() {
	err := zaplog.InitZap()
	if err != nil {
		fmt.Println("日志初始化失败，err:", err)
	}
}
