package channel

import (
	"fmt"
	"time"
)

var c = make(chan string)

func testChannel() {
	time.Sleep(time.Millisecond * 3000)
	c <- "我是通道 channel"
}

// Channel 主协程
func Channel() {
	//开启协程
	go testChannel()

	println("wait")
	//把通道的值赋值给 dataStr
	dataStr := <-c
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", dataStr)
}
