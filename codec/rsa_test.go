package codec

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

// func TestRSA(T *testing.T) {
// 	data, _ := hex.DecodeString("4e53554356318000ad402145b5f6c34ff6d4d112bd2910ccbc68673a38456acd22f7cf6a523efd70e1f40579ed14cd4175d99d4af727c89926e84dcb020a1d96119cb2848033f29b95a150d8fca178bb7b2c6b307e51a1fc8a6162d5730ff7b5b1723f34b0253eae9c521f9b68eb70f45d3d1e257ef636f3bc3cb2c63c371fb75e561e77bcf961ed")
// 	data = data[8:]
// 	data2, err := RsaDecrypt(data[:128], )
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	println(string(data2))
// 	enc, _ := RsaEncrypt([]byte("q342342352352352354"))
// 	fmt.Println(len(enc))
// }

var private = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCjkUQA1TdwPpa3HsEDE9/C3ZRxFTg0gTNkPyoSjHGmNfR5HL1s
j1656xiE8TrW/oI5errdpodsvpsqm4d4uovvV/8Fa70lxfQ1i2nWQN4YNzH2xfni
6ntAP8cXmvRCZKc5NtW6KhMls9397eubiDpIzUYS/gKwZp6uPF3NZkjmWwIDAQAB
AoGADc3lulhETIHLwHqk+XiE6vI+Y+jRjITW7H/0MgqOUOO+1TXaur3C1dgEgrvF
Jn3mSamU+b3jMgdIGylzHnpH8vF70gpzB8H9D/3lnUnzyKNSv4UCNW2bbCU1juQw
S4fWEWjmi/OGNNOEESmnXumM0Rxx+4QTfewR1cqFeGSeKfECQQC5znYgCVLPZM+s
7daEtubZnuvYS/NdO9mRDdwmiKr+Qu8tXk7Bbl6VznvYzkgjhDjy5/LdA3vZ6Uvs
ozxrHbXpAkEA4VwLAHyY1C3UuPc/Bqr75uIC7rXUh7NPJAW8alJVQNPHAb0tYld3
Ycekn6p3ryY5osaLmVYgtr8EhbsUVSqbowJAJCmFnfiSkGCrdpmXfZ7nUQV4G1G0
3LlwP6X16d4BgZjfWfIX29eyOu/D9M85BQiP2N7ByrgJ28BnEXg3oxVWOQJBAKCR
e0lKfX3YddOgXqi6lSbpbBt3NMnHSaEp8Rh0N0gsXIPxrW9/UJE7tSEKTaJfAvvm
qTqEsmRi7671H8Sayi8CQQCL2ScpR6uDB1Xu4kc7MoPdiQoWLvmjvdOZfTw+E8yT
IDDvjkD8dylB4vLKsV4M5MtyMbWvRORcyT+pnfyxFFwE
-----END RSA PRIVATE KEY-----
`)

func TestEncrypt(T *testing.T) {
	a := []byte("{\"K\":\"WAIT\",\"T\":636983611966492353,\"Token\":636983611883881450}")
	b, _ := RsaEncrypt(a)
	fmt.Println(base64.StdEncoding.EncodeToString(b))
	// 从 PEM 私钥导出真正私钥
	block, _ := pem.Decode(private)
	if block == nil {
		println("can't parse private key")
		os.Exit(-1)
	}
	Priv, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
	c, _ := RsaDecrypt(b)
	fmt.Println(string(c))
}
