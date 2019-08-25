package udp

// 这里是 response 拦截器
// 当相应请求被处理之后才交由上次调用函数

import (
	"bytes"
	"encoding/json"
	"strings"

	"../codec"
	"../curl"
	"../system"
	"../types"
)

func parse(cipher []byte) {
	data := cipher[8:]
	// println(string(cipher[:6]))
	if string(cipher[:6]) == "NSUCV1" {
		var err error
		data, err = codec.RsaDecrypt(data)
		if err != nil {
			println("RSA decrypt error")
			return
		}
	}
	s, _ := codec.DecodeGBK(data)
	// println(string(s))
	var j interface{}
	decoder := json.NewDecoder(bytes.NewReader(s))
	decoder.UseNumber()
	decoder.Decode(&j)
	// json.Unmarshal(s, &j)
	update(j.(map[string](interface{})))
}

// 更新信息
func update(res map[string]interface{}) {
	t, has := res["Token"]
	if has {
		newT, _ := t.(json.Number).Int64()
		token = uint64(newT) // 更新 Token
	}
	operate(res)
}

// 从服务器返回的操作，是返回的操作
func operate(res map[string]interface{}) {
	method := res["K"].(string)
	result := res["Result"].(bool)
	switch method {
	case "LOGIN":
		// 公钥验证成功，等待外网网络连通
		if result {
			phase = 2
			Wait()
		} else {
			phase = 4
			println("cert error")
			tokenVerify := strings.SplitN(res["URL"].(string), "?", 1)[1]
			curl.Start(tokenVerify)
			// os.Exit(-2)
		}
	case "WAIT":
		// 开网（外网畅通）成功后
		// if result {
		// 	Keep()
		// }
	case "LOGOUT":
		println("注销成功")
	case "KEEP":
	case "UPDATE":
		println("发现新版本")
	case "MSG":
		println("推送了新的消息")
	case "TIME":
		// println("服务器要求更新时间")
		time, _ := res["T"].(json.Number).Int64()
		tm := types.UnTick(uint64(time))
		system.SetSystemDate(tm)
	default:
		println(method)
		println("服务器发送了特别的请求，注意监测并记录到 log 文件中来反制")
	}
}
