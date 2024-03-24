package cryptoes

import "bytes"

// 可逆加密方法

// 对称加密，加密和解密过程使用的是相同的密钥，因此密钥的安全性至关重要。如果密钥泄露，攻击者可以轻易地破解加密数据。
// 常见的对称加密算法包括DES、3DES、AES等。其中，AES算法是目前使用最广泛的对称加密算法之一，具有比较高的安全性和加密效率。

//不同的加密模式采用不同的策略来处理这些问题。
//ECB（Electronic Codebook）模式： 最简单的模式，将明文分成固定大小的块，然后分别加密。这种模式的问题是相同的明文块会加密成相同的密文块，因此缺乏随机性和安全性。
//CBC（Cipher Block Chaining）模式： 每个明文块与前一个密文块进行异或操作，然后加密。这增加了随机性，并且在错误传播方面具有良好的性质，但需要一个初始化向量（IV）。
//CFB（Cipher Feedback）模式： 通过将前一个密文块作为输入与明文块进行异或操作，产生密文块。这种模式对于错误传播具有良好的性质，但也需要一个初始化向量。
//OFB（Output Feedback）模式： 类似于 CFB，但是通过生成一个密钥流来与明文进行异或操作。也对错误传播具有良好的性质。
//CTR（Counter）模式： 使用计数器来生成一个密钥流，然后与明文进行异或操作。在并行计算环境中效率较高。

//ECB是最快、最简单的分组密码模式，但它的安全性最弱，一般不推荐使用ECB加密消息，但如果是加密随机数据，如密钥，ECB则是最好的选择。
//CBC适合文件加密，而且有少量错误时不会造成同步失败，是软件加密的最好选择。
//CFB通常是加密分组序列所选择的模式，它也能容忍少量错误扩展，且具有同步恢复功能。推荐使用CTR模式代替。
//OFB是在极易出错的环境中选用的模式，但需有高速同步机制。推荐使用CTR模式代替。

// 不足字符需要填充
func PaddingText(plainText []byte, blockSize int) []byte {
	padNum := blockSize - len(plainText)%blockSize
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)
	plainText = append(plainText, newPlain...)
	return plainText
}

func UnPaddingText(plainText []byte) []byte {
	length := len(plainText)
	number := int(plainText[length-1])
	return plainText[:length-number]
}
