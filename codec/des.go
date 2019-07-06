package codec

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
)

// DesEncrypt DES 加密
func DesEncrypt(data []byte, password string) []byte {
	key := genDESKey(password)
	block, _ := des.NewCipher(key)         // 从密钥创建真正的密钥块，由于保证 key 一定是 8 字节，省略错误
	data = pkcs5Padding(data)              // 明文填充
	cryptedData := make([]byte, len(data)) // 创建存放密文的容器
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(cryptedData, data)
	return cryptedData
}

// DesDecrypt DES 解密
func DesDecrypt(data []byte, password string) []byte {
	key := genDESKey(password)
	block, _ := des.NewCipher(key)         // 从密钥创建真正的密钥块，由于保证密钥一定是 8 字节，省略错误
	cryptedData := make([]byte, len(data)) // 创建存放密文的容器
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(cryptedData, data)
	return pkcs5UnPadding(cryptedData)
}

func genDESKey(password string) []byte {
	/* 密钥生成 */
	pw := []byte(password)    // 将密码转换为字节码
	hash := md5.Sum(pw)       // 计算密码的 MD5 值
	key := make([]byte, 8)    // 创建存储密钥的 slice
	hex.Encode(key, hash[:4]) // 将 MD5 的前 4 个字节编成 8 位 hex 码
	key = bytes.ToUpper(key)  // hex 码转大写
	return key
}

func pkcs5Padding(plain []byte) []byte {
	paddingSize := 8 - len(plain)%8 // 需要填充的字节数（DES 中为 8 字节）
	padding := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(plain, padding...)
}

func pkcs5UnPadding(plain []byte) []byte {
	length := len(plain)
	upPaddingSize := int(plain[length-1])
	return plain[:length-upPaddingSize]
}
