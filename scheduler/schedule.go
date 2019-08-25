package scheduler

type info struct {
	operation, value string
}

var queue = make(chan *info, 8)
var state = &info{operation: "stop"}

func createResource(operation, value string) *info {
	s := &info{operation, value}
	return s
}

func release() {

}

func run() {
	m := map[string]int{"stop": 0, "start": 1, "verify": 2, "wait": 3, "keep": 4}
	for msg := range queue {
		code := m[msg.operation]
		if code > 0 && code >= m[state.operation] {
			state = msg
		} else if code <= 0 {
			// 将 udp 发送队列中的元素全部清空
			//
		}
	}
	// start -> (verify) -> wait -> keep
	// ==> stop 任意状态均可以到 stop 状态转换
}
