package main

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	// 创建验证码实例
	cap := captcha.New()

	// 设置字体
	err := cap.AddFont("other/comic.ttf")
	if err != nil {
		log.Fatal(err)
	}
	// 设置验证码图片尺寸
	cap.SetSize(100, 40)
	// 设置随机背景颜色
	cap.SetBkgColor(color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: uint8(255),
	})

	// 添加干扰线
	cap.SetDisturbance(captcha.HIGH)

	// 生成验证码图片和对应的字符串
	img, str := cap.Create(4, captcha.LOWER)

	// 保存图片到文件
	file, _ := os.Create("captcha.png")
	defer file.Close()
	png.Encode(file, img)

	// 输出验证码字符串
	fmt.Println("Captcha string:", str)
}

//func main() {
//	captchaID := captcha.New()
//	outfile, err := os.Create("captcha.png")
//	if err != nil {
//		log.Fatal(err)
//	}
//	_ = captcha.WriteImage(outfile, captchaID, 200, 100)
//}

//func main() {
//	// 配置验证码生成器
//	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
//
//	// 生成验证码实例
//	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
//
//	// 生成图片验证码
//	id, b64s, _, err := captcha.Generate()
//	if err != nil {
//		fmt.Println("Failed to generate captcha:", err)
//		return
//	}
//
//	fmt.Println("Captcha ID:", id)
//
//	// 将base64编码的图片数据解码
//	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64s))
//	m, _, err := image.Decode(reader)
//	if err != nil {
//		fmt.Println("Failed to decode captcha image:", err)
//		return
//	}
//
//	// 保存图片到文件
//	outfile, err := os.Create("captcha.png")
//	if err != nil {
//		fmt.Println("Failed to create output file:", err)
//		return
//	}
//	defer outfile.Close()
//
//	err = png.Encode(outfile, m)
//	if err != nil {
//		fmt.Println("Failed to save captcha image:", err)
//		return
//	}
//
//	fmt.Println("Captcha image saved as captcha.png")
//}
