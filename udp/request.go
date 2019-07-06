package udp

import (
	"encoding/json"
	"time"

	"../codec"
	"../types"
)

var loginInfo *types.LoginInfo // 记录登录信息
var queue = make(chan []byte)
var token uint64
var waitTicker *time.Ticker

// Login 登录发送请求
func Login() {
	ticker := time.NewTicker(time.Second * 10)
	for range ticker.C {
		loginInfo.T = types.Tick() // 更新最新的时间戳
		info, _ := json.Marshal(loginInfo)
		enQueue(info)
	}
}

// Wait 等待开网（点击登录之后） 30 秒内循环请求 30 次
func Wait() {
	waitTicker = time.NewTicker(time.Second)
	waitOrKeep(waitTicker, true)
}

// Keep 保持连接 10 秒发送一次
func Keep() {
	// 结构和等待开网相同，只是发送频率不同
	keepTicker := time.NewTicker(time.Second * 10)
	waitOrKeep(keepTicker, false)
}

func waitOrKeep(ticker *time.Ticker, wait bool) {
	cnt := 0
	for range ticker.C {
		if wait {
			cnt++
			if cnt > 30 {
				ticker.Stop()
			}
		}
		keepInfo := types.KeepInfo{K: "Wait", T: types.Tick(), Token: token}
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
