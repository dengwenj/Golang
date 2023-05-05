### golang并发编程之协程

* golang 中的并发是**函数**相互独立运行的能力。Gotoutines（协程） 是并发运行的函数。golang提供了 Gotoutines 作为并发处理操作的一种方式。

* 创建一个协程非常简单，就是在一个任务函数前面加一个 go 关键字

* ```go
  go task()
  ```

```go
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
```

