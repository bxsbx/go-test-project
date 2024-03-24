package cryptoes

import (
	"crypto/aes"
	"crypto/cipher"
)

//AES 支持 5 种加密模式。 (ECB | CBC | CFB | OFB | CTR) 支持五种模式

//AES（Advanced Encryption Standard）即高级加密标准，是一种对称加密算法，被广泛应用于数据加密和保护领域。
//AES算法使用的密钥长度为128位、192位或256位，比DES算法的密钥长度更长，安全性更高。

//AES算法采用的密钥长度更长，密钥空间更大，安全性更高，能够有效地抵抗暴力破解攻击。
//当然，因为密钥长度较长，需要的存储也更多。
//对于对称加密算法而言，最大的痛点就在于密钥管理困难，相比而言。

// AES加密
func EncryptAES(src, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	// 填充数据
	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// AES解密
func DecryptAES(src, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockeMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockeMode.CryptBlocks(dst, src)
	//去掉填充
	return UnPaddingText(dst)
}
