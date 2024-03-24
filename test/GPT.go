package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/url"
	"path"
)

func RSAGenerateKey(bits int) (privateKeyStr, publicKeyStr string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	// 私钥
	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{
		Type:  "RES Private Key ",
		Bytes: privateBytes,
	}
	privateWriter := bytes.Buffer{}
	err = pem.Encode(&privateWriter, &privateBlock)
	if err != nil {
		return
	}
	privateKeyStr = privateWriter.String()

	// 公钥
	publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}
	publicBlock := pem.Block{
		Type:  "RSA Public Key",
		Bytes: publicBytes,
	}
	publicWriter := bytes.Buffer{}
	err = pem.Encode(&publicWriter, &publicBlock)
	if err != nil {
		return
	}
	publicKeyStr = publicWriter.String()
	return
}

// 加密
func EncryptRsa(src []byte, privateKeyStr string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := pub.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)

}

// 解密
func Decrypt(src []byte, publicKeyStr string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)

}

func main() {
	u, err := url.Parse("http://example.com/path/to/resource?param1=value1&param2=value2")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	u.RawQuery = "" // 清空查询参数
	u.Fragment = "" // 清空fragment

	cleanPath := path.Clean(u.Path) // 清理路径
	u.Path = cleanPath

	fmt.Println("URL without parameters:", u.String())
}
