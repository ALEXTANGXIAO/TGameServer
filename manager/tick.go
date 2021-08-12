package manager

import (
	"fmt"
	"time"
)

type Ticker struct {
	TickTime float64
	Func     interface{}
}

func new() {
	ticker := time.NewTicker(time.Second) //定时1秒 如果想间隔2秒执行一次只需将1改成2
	go func() {
		for t := range ticker.C { //定时干活
			fmt.Println("tick", t) //干活
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop() //结束工作
}
