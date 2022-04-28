package trick

import (
	"fmt"
	"time"
)

func DemoCron() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("ticker ticker ticker ...")
	}
}

// 定义函数类型
type Fn func()

// 定时器中的成员
type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}

func NewMyTick(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

// 启动定时器需要执行的任务
func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner()
		}
	}
}
