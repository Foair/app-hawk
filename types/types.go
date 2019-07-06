package types

// LoginInfo 登陆所需要的信息
type LoginInfo struct {
	UserID, UserPW, IP, MAC, ADPID,
	Port, CName, OSVer, CVer, RSA, MD5, K string
	T uint64
}

// KeepInfo 保持连接需要的信息
type KeepInfo struct {
	K        string
	T, Token uint64
}

// UserState 用户状态信息
type UserState struct {
	Password, PubKeyXML, PrivateKey string
}
