package util

import "fmt"

const (
	RED       = "\033[31m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	GREEN     = "\033[32m"
	CYAN      = "\033[36m"
	TURQUOISE = "\033[96m" //蓝绿色
	PURPLE    = "\033[35m"
	RESET     = "\033[0m"
)

func ColorPrint(str interface{}, color string) string {
	var newStr string
	switch color {
	case BLUE, RED, YELLOW, GREEN, CYAN, TURQUOISE, PURPLE:
		newStr = fmt.Sprintf("%s%v%s\n", color, str, RESET)
	default:
		newStr = fmt.Sprintf("%v\n", str)
	}
	return newStr
}
