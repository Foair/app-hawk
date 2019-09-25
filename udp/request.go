package udp

import (
	"encoding/json"
	"sync"
	"time"

	"why.moe/hawk/codec"
	"why.moe/hawk/types"
)

var loginInfo *types.LoginInfo // 记录登录信息
var queue = make(chan []byte)
var token uint64
var waitTicker *time.Ticker
var phase = 0 // 0 停止 1 Login() 运行中 2 Wait() 中 3 Keep() 中 4 错误
var lock sync.Mutex

// Login 登录发送请求
func Login() {
	phase = 1
	ticker := time.NewTicker(time.Second * 4)
	for range ticker.C {
		if phase >= 2 {
			ticker.Stop()
			Wait()
			// return
		}
		loginInfo.T = types.Tick() // 更新最新的时间戳
		info, _ := json.Marshal(loginInfo)
		enQueue(info)
	}
}

// Wait 等待开网（点击登录之后） 30 秒内循环请求 30 次
func Wait() {
	lock.Lock()
	if phase == 3 {
		return
	}
	phase = 3
	lock.Unlock()
	waitTicker = time.NewTicker(time.Second * 10)
	waitOrKeep(waitTicker, true)
}

// Keep 保持连接 10 秒发送一次
func Keep() {
	// 结构和等待开网相同，只是发送频率不同
	keepTicker := time.NewTicker(time.Second * 10)
	waitOrKeep(keepTicker, false)
}

func waitOrKeep(ticker *time.Ticker, wait bool) {
	// cnt := 0
	for range ticker.C {
		// if wait {
		// 	cnt++
		// 	if cnt > 30 {
		// 		ticker.Stop()
		// 	}
		// }
		keepInfo := types.KeepInfo{K: "KEEP", T: types.Tick(), Token: token}
		info, _ := json.Marshal(keepInfo)
		enQueue(info)
	}
}

func sendCommon() {

}

func enQueue(data []byte) {
	enc, _ := codec.RsaEncrypt(data)
	// enc := data
	queue <- paddingHeaderSign(enc)
}
