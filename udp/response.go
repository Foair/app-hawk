package udp

// 这里是 response 拦截器
// 当相应请求被处理之后才交由上次调用函数

import "../codec"

func parse(cipher []byte) {
	data := cipher[8:]
	// println(string(cipher[:6]))
	if string(cipher[:6]) == "NSUCV1" {
		var err error
		data, err = codec.RsaDecrypt(data)
		if err != nil {
			return
		}
	}
	// println(string(data))
}

// 更新信息
func update(res map[string]string) {
	token, has := res["Token"]
	if has {
		println(token)
		println("更新 token")
	}
}

// 从服务器返回的操作，是返回的操作
func oprate(method string) {
	switch method {
	case "LOGIN":
		// 公钥验证成功，等待外网网络连通
		Wait()
	case "WAIT":
		// 开网（外网畅通）成功后
		waitTicker.Stop()
		Keep()
	case "LOGOUT":
		println("注销成功")
	case "KEEP":
		println("连接失效")
	case "UPDATE":
		println("发现新版本")
	case "MSG":
		println("推送了新的消息")
	default:
		println("服务器发送了特别的请求，注意监测并记录到 log 文件中来反制")
	}
}
