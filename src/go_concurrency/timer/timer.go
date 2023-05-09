package timer

import (
	"fmt"
	"time"
)

func LTimer() {
	//timer := time.NewTimer(time.Second * 2)
	//
	//fmt.Printf("%v\n", time.Now())
	//t := <-timer.C // 阻塞，等2秒结束再执行
	//fmt.Printf("%v\n", t)
	//fmt.Printf("%v\n", time.Now())

	//<-time.After(time.Second * 2)
	//println("等两秒后再执行")

	//timer := time.NewTimer(time.Second * 2)
	//go func() {
	//	<-timer.C
	//	println("stop 了就不会执行了")
	//}()
	//stop := timer.Stop()
	//if stop {
	//	println("不会执行 上面协程 里面了，会执行这里")
	//}

	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("%v\n", timer)
	//timer.Reset(time.Second)
	<-timer.C
	fmt.Printf("%v\n", "等等多少秒")
}
