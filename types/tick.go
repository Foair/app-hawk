package types

import "time"

// Tick 获得当前 C# 时间
func Tick() uint64 {
	now := time.Now()
	tick := now.UnixNano()/100 + 621356256000000000
	return uint64(tick)
}
