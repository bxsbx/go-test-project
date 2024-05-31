package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// 要生成的文本信息
	text := "Hello, this is a QR code!"

	// 创建 QR 码的编码器
	qrCode, _ := qr.Encode(text, qr.M, qr.Auto)

	// 将编码后的 QR 码转换为图片
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// 创建输出文件
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// 将编码后的 QR 码写入文件
	png.Encode(file, qrCode)
}
