package config

import "net"

const (
	// ClientIP UDP 监听 IP
	ClientIP = "127.0.0.1"
	// ServerIP UDP 服务端 IP
	ServerIP = "118.113.89.40"
	// FakeIP 伪造 IP
	FakeIP = ""

	// MAC MAC 地址
	MAC = "2076933E40D3"
	// ComputerName 计算机名
	ComputerName = "DESKTOP-B0DPOT1"
	// UserID 用户名
	UserID = "16310120537"
	// AdpID 适配器 ID
	AdpID = "{BE34365E-52CB-49D6-B2D3-355ED833FD5F}"
	// OSVer 操作系统版本
	OSVer = "Microsoft Windows NT 6.2.9200.0"
	// CCVer CC 版本号
	CCVer = "1.16.3.8013"
	// CCMD5 CC MD5 校验
	CCMD5 = "DB2F9160F56FC8AFD7AA12B8DDC8DA23"
)

var (
	// SrcEP 客户端 End Point
	SrcEP = &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	// DstEP 服务端 End Point
	DstEP = &net.UDPAddr{IP: net.ParseIP(ServerIP), Port: 8080}
)
