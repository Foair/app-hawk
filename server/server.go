package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("../user.json")
	_, _ = w.Write(file)
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/query", query)
	_ = http.ListenAndServe(":8848", nil)
}

func query(w http.ResponseWriter, req *http.Request) {
	status := "s"
	ip := "10.0.2.2"
	port := "8888"
	_, _ = w.Write([]byte(status + "$" + ip + "$" + port))
}

func cord(w http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("cord.log")
	// 每次启动程序的时候，将 cord.log 重命名为 cord + 时间戳 + .txt 的形式（在启动脚本里面完成）
	// 保证最新的文件名都是 cord.log
	_, _ = w.Write(file)
}

func stop(w http.ResponseWriter, req *http.Request) {
	// 停止所有请求的发送，和对回应的响应
}

func reset(w http.ResponseWriter, req *http.Request) {
	// 软重启
}

func reload(w http.ResponseWriter, req *http.Request) {
	// 硬重启（通过脚本）
	runShell("./reload.sh")
}

func verify(w http.ResponseWriter, req *http.Request) {
	// 将所有的响应回归到验证证书的状态
}

// 清理往期日志（使用脚本）
func clean(w http.ResponseWriter, req *http.Request) {
	runShell("./cleanup.sh")
}

func runShell(filename string) {
	cmd := exec.Command(filename)
	_ = cmd.Run()
}
