package goroutines

import "time"

func testGoroutines(lang string) {
	for i := 0; i < 5; i++ {
		println(lang)
		time.Sleep(time.Millisecond * 200)
	}
}

// Goroutines 协程
func Goroutines() {
	go testGoroutines("golang") // go 启动了一个协程来执行
	go testGoroutines("js")     // 2
	time.Sleep(time.Millisecond * 2000)
	println("主协程") // 3 主函数退出，程序就结束了，协程也就结束了
}
