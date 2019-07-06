package codec

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 公钥和私钥可以从文件中读取

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC2VZzq6r1Z+GhGp1mOULg/zL8E
DAKW2k5q4U+FN4OiK9TJapjmoOYD7AGQOrBZxTUobyiMwMT0wlhD9Dulp1ClhYyg
asK9I4n6P+PszpuLyppEqvGiVNBvDdyUVsj6TWBNN9qiNjAxt4W2h/oAXIBuemwC
5f6jS/r1nAR6Y8jkNwIDAQAB
-----END PUBLIC KEY-----
`)

// Priv 私钥
var Priv *rsa.PrivateKey

// RsaEncrypt 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		println("abc")
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		println("abc")
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	round := len(origData) / 117
	if len(origData)%117 != 0 {
		round++
	}
	var buffer []byte
	for i := 0; i < round; i++ {
		var d []byte
		if i == round-1 && 117*(i+1) > len(origData) {
			d, _ = rsa.EncryptPKCS1v15(rand.Reader, pub, origData[117*i:len(origData)])
		} else {
			d, _ = rsa.EncryptPKCS1v15(rand.Reader, pub, origData[117*i:117*(i+1)])
		}
		buffer = append(buffer, d...)
	}

	return buffer, nil
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, Priv, ciphertext)
}
