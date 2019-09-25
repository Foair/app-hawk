package curl

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"why.moe/hawk/codec"
)

type res struct {
	r, key string
}

func Start(token string) {
	client := &http.Client{Transport: transport}
	cookieStore, _ := cookiejar.New(nil)
	if token != "" {
		profile := "https://net.nsu.edu.cn/nsucweb/userinfo.aspx?" + token
		req(client, profile, cookieStore)
	}
	ticker := time.NewTicker(time.Second * 10)
	for range ticker.C {
		println(">>>")
		t := &res{}
		r := req(client, "https://net.nsu.edu.cn/nsucweb/verifyphone.aspx?IsMobile=false", cookieStore)
		json.Unmarshal(r, t)
		if t.r == "1" {
			ticker.Stop()
			encoded, _ := base64.StdEncoding.DecodeString(t.key)
			println(encoded)
			codec.Priv = codec.XMLToPrivateKey(string(encoded))
			os.Exit(2)
		}
	}
}
