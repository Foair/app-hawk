package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"os"

	"why.moe/hawk/codec"
	"why.moe/hawk/config"
	"why.moe/hawk/udp"
)

func main() {
	initial()
	udp.LinkStart(password, rsaPublicKey)
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
}

func initial() {
	// 读取用户状态 JSON
	userInfo, _ := ioutil.ReadFile(os.Args[2])

	// 初始化全局状态
	_ = json.Unmarshal(userInfo, &config.Dict)
	password = config.Dict.Password
	rsaPublicKey = config.Dict.PubKeyXML

	// 从 PEM 私钥导出真正私钥
	block, _ := pem.Decode([]byte(config.Dict.PrivateKey))
	if block == nil {
		println("can't parse private key")
		os.Exit(-1)
	}
	rsaPrivateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
	codec.Priv = rsaPrivateKey
}
