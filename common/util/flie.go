package util

import "os"

func WriteToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		file, err = os.Create(filePath)
	}
	if err != nil {
		return err
	}
	defer file.Close()

	// 将内容写入文件
	_, err = file.WriteString(content + "\n")
	return err
}
