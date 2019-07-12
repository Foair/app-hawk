package types

import "time"

// Tick 获得当前 C# 时间
func Tick() uint64 {
	now := time.Now()
	tick := now.UnixNano()/100 + 621356256000000000
	return uint64(tick)
}

// UnTick 从 C# 时间还原出时间戳
func UnTick(tick uint64) uint64 {
	timeStamp := tick - 621356256000000000
	timeStamp /= 10000000
	return timeStamp
}
