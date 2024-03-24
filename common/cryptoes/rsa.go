package cryptoes

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

//非对称加密算法

//非对称加密算法需要两个密钥，这两个密钥互不相同，但是相互匹配，一个称为公钥，另一个称为私钥。
//使用其中的一个加密，则使用另一个进行解密。例如使用公钥加密，则需要使用私钥解密。

//RSA算法的优点是安全性高，公钥可以公开，私钥必须保密，保证了数据的安全性；可用于数字签名、密钥协商等多种应用场景。
//缺点是加密、解密速度较慢，密钥长度越长，加密、解密时间越长；密钥长度过短容易被暴力破解，密钥长度过长则会增加计算量和存储空间的开销。

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
func DecryptRsa(src []byte, publicKeyStr string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)

}
