package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	filePath := "C:\\Users\\WangYu\\Desktop\\新建文本文档.txt" // 替换为你的文件路径

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件时出错:", err)
		return
	}
	defer file.Close()

	// 读取文件内容的前 512 个字节，用于判断 MIME 类型
	buffer := make([]byte, 2048)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("读取文件内容时出错:", err)
		return
	}

	// 获取文件的 MIME 类型
	mimeType := http.DetectContentType(buffer)
	fmt.Println("文件的 MIME 类型为:", mimeType)
}
