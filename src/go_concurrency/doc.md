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

### 并发编程之通道（channel）

* golang 提供了一种称为通道的机制，用于在 goroutine （协程）之间**共享数据**。当作为 goroutine执行并发活动时，需要在 goroutine 之间共享资源或数据，通道充当 goroutine 之间的管道并提供一种机制来保证**同步**交换。

* 需要再声明通道时指定数据类型。我们可以共享内置、命名、结构体和引用类型的值和指针。数据在通道上传递，在任何给定时间只有一个 goroutine 可以访问数据项，因此按照设计不会发生数据竞争。

* 根据数据交换的行为，有两种类型的通道，无缓冲通道和缓冲通道。无缓冲通道用于执行 goroutine 之间的同步通信，而缓冲通道用于执行异步通信。无缓冲通道保证在发送和接收发生的瞬间执行两个 goroutine 之间的交换。缓冲通道没有这样的保证。

* 通道由 make 函数创建，该函数指定 chan 关键字和通道的元素类型。

* ```go
  unBuffered := make(chan int) // 整型无缓冲通道
  buffered := make(chan int, 10) // 整型有缓冲通道
  ```

* 使用内置函数 `make`创建无缓冲和缓冲通道。`make`的第一个参数需要关键字 `chan`，然后是通道允许交换的数据类型。

* 这是将值发送到通道的代码块需要使用 `<-`运算符

* ```go
  goroutine1 := make(chan string, 5) // 字符串缓冲通道
  goroutine1 <- "你好啊" // 通过通道发送字符串
  ```

* 一个包含 5 个值的缓冲区的字符串类型的 goroutine1 通道。然后我们通过通道发送字符串。

* 这是从通道接收值的代码块：

* ```go
  data := <-goroutine1 // 把通道的值赋值给 data
  ```

* <- 运算符附加到通道变量（goroutine）的左侧，以接收来自通道的值

* **无缓冲通道**：在无缓冲通道中，在接收到任何值之前没有能力保存它，在这种类型的通道中，发送和接收 goroutine 在任何发送或接收操作完成之前的同一时刻都准备就绪。如果两个 goroutine 没有在同一时刻准备好，则通道会让执行其各自发送或接收操作的 goroutine 首先等待。同步是通道上发送和接收之间交互的基础，没有另一个就不可能发生。

* **缓冲通道**：在缓冲通道中，有能力在接收到一个或多个值之前保存他们。在这种类型的通道中，不要强制 goroutine 在同一时刻准备好执行发送和接收。当发送或接收阻塞时也有不同的条件。只有当通道中没有要接收的值时，接收才会阻塞。仅当没有可用缓冲区来放置正在发送的值时，发送才会阻塞。

* 通道的发送和接收特性

* 1、对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。

* 2、发送操作和接收操作中对元素值的处理都是不可分割的。

* 3、发送操作在完全完成之前会被阻塞，接收操作也是如此。

```go
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
```

### 并发编程之 WaitGroup 实现协程之间（都执行）同步

* 查看添加 WaitGroup 和不添加 WaitGroup 的区别

```go
package waitgroup

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func testWaitGroup(i int) {
	fmt.Printf("%v\n", i)
	defer wg.Done() // wg.Add(-1)
}

func WaitGroup() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个 goroutine 就登记 +1
		go testWaitGroup(i)
	}

	// 等待，没有添加过就执行了。阻塞了同步。等待所有的 goroutine都结束
	wg.Wait()
	println("end...")
}
```

### 并发编程之使用 runtime 里关于协程 api

```go
package runtimedwj

import (
	"fmt"
	"runtime"
	"time"
)

func testRunTime(msg string) {
	for i := 0; i < 2; i++ {
		println(msg)
	}
}

func testRunTime1() {
	for i := 0; i < 10; i++ {
		println(i)
		if i >= 3 {
			// 退出当前协程
			runtime.Goexit()
		}
	}
}

func testRunTime2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("testRunTime2: %v\n", i)
	}
}

func testRunTime3() {
	for i := 0; i < 10; i++ {
		fmt.Printf("testRunTime3: %v\n", i)
	}
}

func RunTime() {
	//go testRunTime("JavaScript") // 子协程
	//runtime.Gosched()            // 我有权利执行任务，让其他子协程来执行
	//println("end...")

	//	退出当前 协程
	//go testRunTime1()
	//time.Sleep(time.Second)

	fmt.Printf("%v\n", runtime.NumCPU())
	//	设置 cpu 核
	runtime.GOMAXPROCS(1)
	go testRunTime2()
	go testRunTime3()
	time.Sleep(time.Second)
}
```

