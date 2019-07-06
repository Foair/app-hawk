package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"os"

	"./codec"
	"./types"
	"./udp"
)

func main() {
	initail()
	udp.LinkStart(password, rsaPublicKey, rsaPrivateKey)
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func initail() {
	// 读取用户状态 JSON
	var dict types.UserState
	userInfo, _ := ioutil.ReadFile("user.json")

	// 初始化全局状态
	json.Unmarshal(userInfo, &dict)
	password = dict.Password
	rsaPublicKey = dict.PubKeyXML

	// 从 PEM 私钥导出真正私钥
	block, _ := pem.Decode([]byte(dict.PrivateKey))
	if block == nil {
		println("can't parse private key")
		os.Exit(-1)
	}
	rsaPrivateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
	codec.Priv = rsaPrivateKey
}
