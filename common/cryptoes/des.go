package cryptoes

import (
	"crypto/cipher"
	"crypto/des"
)

//加密过程
//使用des.NewCipher获取块
//填充明文
//使用cipher.NewCBCEncrypter生成CBC模式块
//使用模式块的CryptBlocks进行加密，由于可指向同一内存地址，我们仍使用同一变量，节省内存

// des (ECB | CBC | CFB | OFB | CTR) 支持五种模式
// key 的值为8位，否则报错
// DES加密
func EncryptDES(src, iv, key []byte) []byte {
	block, _ := des.NewCipher(key)
	// 填充数据
	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// DES解密
func DecryptDES(src, iv, key []byte) []byte {
	block, _ := des.NewCipher(key)
	blockeMode := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(src))
	blockeMode.CryptBlocks(dst, src)
	//去掉填充
	return UnPaddingText(dst)
}

func Encrypt3DES(src, key []byte) []byte {
	//des包下的三次加密接口
	block, _ := des.NewTripleDESCipher(key)
	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

func Decrypt3DES(src, key []byte) []byte {
	block, _ := des.NewTripleDESCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return UnPaddingText(dst)
}
