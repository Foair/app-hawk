package udp

import (
	"crypto/rsa"
	"net"
	"strings"

	"../config"
	"../types"
)

// LinkStart 开启 UDP
func LinkStart(password, rsaPublicKey string, rsaPrivateKey *rsa.PrivateKey) {
	conn, _ := net.DialUDP("udp4", config.SrcEP, config.DstEP)
	EP := strings.SplitN(conn.LocalAddr().String(), ":", 2)
	ipAddr, port := EP[0], EP[1]
	println(ipAddr, port)
	// defer conn.Close()

	loginInfo = &types.LoginInfo{
		UserID: config.UserID,
		UserPW: password,
		IP:     "10.98.38.92",
		MAC:    config.MAC,
		ADPID:  config.AdpID,
		T:      types.Tick(),
		Port:   port,
		CName:  config.ComputerName,
		OSVer:  config.OSVer,
		RSA:    rsaPublicKey,
		MD5:    config.CCMD5,
		K:      "LOGIN",
		CVer:   config.CCVer}
	go sender(conn)
	go receiver(conn)
	Login()
}

// 接收进程
func receiver(conn *net.UDPConn) {
	data := make([]byte, 1024) // 接收缓存
	for {
		conn.Read(data) // 收到的字节数
		println("<<<")
		parse(data)
	}
}

// 发送进程
func sender(conn *net.UDPConn) {
	for data := range queue {
		// println(">>>")
		conn.Write(data)
	}
}
