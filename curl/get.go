package curl

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"../config"
)

var reqEP *net.TCPAddr
var transport *http.Transport

// init 支持指定本地 IP 访问，实现单线多拨
func init() {
	reqEP = &net.TCPAddr{IP: net.ParseIP(config.ClientIP)}
	transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			LocalAddr: reqEP,
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func req(client *http.Client, url string, cookieStore *cookiejar.Jar) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", config.UserAgent)
	client.Jar = cookieStore

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	return body
}

func main() {
	client := &http.Client{Transport: transport}
	cookieStore, _ := cookiejar.New(nil)
	fmt.Println(cookieStore)
	req(client, "https://httpbin.org/cookies/set?a=b", cookieStore)
	fmt.Println(cookieStore)
	req(client, "https://httpbin.org/cookies", cookieStore)
}
