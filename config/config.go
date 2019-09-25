package config

import (
	"net"

	"why.moe/hawk/types"
)

const (
	// ServerIP UDP 服务端 IP
	ServerIP = "118.113.89.40"
	// UserAgent 用户代理
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; Tablet PC 2.0; rv:11.0) like Gecko"

	// OSVer 操作系统版本
	OSVer = "Microsoft Windows NT 6.2.9200.0"
	// CCVer CC 版本号
	CCVer = "1.16.3.8013"
	// CCMD5 CC MD5 校验
	CCMD5 = "DB2F9160F56FC8AFD7AA12B8DDC8DA23"
)

var (
	// SrcEP 客户端 End Point
	SrcEP = &net.UDPAddr{IP: net.ParseIP(ClientIP), Port: 0}
	// DstEP 服务端 End Point
	DstEP = &net.UDPAddr{IP: net.ParseIP(ServerIP), Port: 8080}
)

// Dict 用于存放用户信息
var Dict types.UserState
