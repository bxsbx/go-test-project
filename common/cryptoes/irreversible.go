package cryptoes

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

//不可逆加密方法

//可加盐再加密

// md5加密(可对字符串加盐，自行处理)
// MD5算法的输出长度为128位，通常用32个16进制数表示。
// MD5有一些优点，比如计算速度快、输出长度固定、应用广泛
// 缺点不安全，可以通过暴力破解或彩虹表攻击等方式，找到与原始数据相同的散列值，从而破解数据
func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func MD5Salt(s string, salt string) string {
	hash := hmac.New(md5.New, []byte(salt))
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

//SHA有三个版本
//散列值长度更长：例如SHA-256算法的散列值长度为256位，而MD5算法的散列值长度为128位，这就提高了攻击者暴力破解或者彩虹表攻击的难度。
//更强的碰撞抗性：SHA算法采用了更复杂的运算过程和更多的轮次，使得攻击者更难以通过预计算或巧合找到碰撞。
//当然，SHA-2也不是绝对安全的，散列算法都有被暴力破解或者彩虹表攻击的风险，所以，在实际的应用中，加盐还是必不可少的。

// SHA-1系列存在缺陷，已经不再被推荐使用
func SHA1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA1Salt(s string, salt string) string {
	hash := hmac.New(sha1.New, []byte(salt))
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA-2算法包括SHA-224、SHA-256,SHA-384、SHA-512散列函数，推荐使用
func SHA2(s string) string {
	hash := sha256.New()
	//hash := sha256.New224()
	//hash := sha512.New()
	//hash := sha512.New384()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA2Salt(s string, salt string) string {
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA-3算法
func SHA3(s string) string {
	hash := sha3.New256()
	//hash := sha3.New224()
	//hash := sha3.New384()
	//hash := sha3.New512()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA3Salt(s string, salt string) string {
	hash := hmac.New(sha3.New256, []byte(salt))
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
