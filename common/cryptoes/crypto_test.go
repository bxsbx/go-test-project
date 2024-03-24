package cryptoes

import (
	"fmt"
	"log"
	"testing"
)

// 测试方法
func TestDEC(t *testing.T) {
	src := []byte("23233134vevrernv")
	iv := []byte("ejfeofjw")
	key := []byte("23999211")
	des := EncryptDES(src, iv, key)
	fmt.Println(string(des))
	text := DecryptDES(des, iv, key)
	fmt.Println(string(text))
}

func Test3DEC(t *testing.T) {
	src := []byte("23233134vevrernv")
	key := []byte("239992112399921123999211")
	des := Encrypt3DES(src, key)
	fmt.Println(string(des))
	text := Decrypt3DES(des, key)
	fmt.Println(string(text))
}

func TestAEC(t *testing.T) {
	src := []byte("23233134vevrernvcwvevwevwevwevwvbrrbrb")
	//key := []byte("23999211239992112399921123999211")
	key := []byte("2399921123999211")
	des := EncryptAES(src, key)
	fmt.Println(string(des))
	text := DecryptAES(des, key)
	fmt.Println(string(text))
}

func TestRSA(t *testing.T) {
	src := []byte("aaaaanvpwwwwwwwwwwpq")
	prikey, pubkey, err := RSAGenerateKey(256)
	if err != nil {
		log.Fatal(err)
	}
	cryptRsa, err := EncryptRsa(src, pubkey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("密文===", fmt.Sprintf("%x", cryptRsa))
	crypt, err := DecryptRsa(cryptRsa, prikey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("明文===", fmt.Sprintf("%s", crypt))
}
