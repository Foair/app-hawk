package system

import (
	"os/exec"
	"strconv"
)

// SetSystemDate 根据 Unix 时间戳设置系统时间
func SetSystemDate(timeStamp uint64) {
	cmd := exec.Command("date", "@"+strconv.FormatUint(timeStamp, 10))
	println("date +%s -s @" + strconv.FormatUint(timeStamp, 10))
	cmd.Run()
}
