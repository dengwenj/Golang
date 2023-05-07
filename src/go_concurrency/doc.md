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

### golang 并发编程之 Mutex 互斥锁实现同步

* 除了使用 channel 实现同步之外，还可以使用 Mutex 互斥锁的方式实现同步

```go
package mutex

import (
	"fmt"
	"sync"
)

var num = 50

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock() // 加锁  别的协程执行不了，只有解了才可以执行
	num++
	fmt.Printf("i--%v\n", num)
	lock.Unlock() // 解锁
}

func sub() {
	defer wg.Done()
	lock.Lock()
	num--
	fmt.Printf("i--%v\n", num)
	lock.Unlock()
}

func LMutex() {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}

	wg.Wait()

	fmt.Printf("主协程:%v\n", num)
}
```

### channel 进行遍历

* 如果通道关闭，读多写少，没有了就是默认值，例如：int 就是0，如果没有关闭就会死锁

```go
gopackage channel

func LChannelReadWrite() {
	// 通道遍历

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 写入通道
		}
		// 通道写完了过后要关闭 close()
		close(c)
	}()

	//	读取通道里面的值
	//cData := <-c
	//fmt.Printf("cData%v\n", cData)
	//遍历读取
	for i := range c {
		num := i
		println(num)
	}

	//for {
	//	data, ok := <-c
	//	if ok {
	//		fmt.Printf("%v\n", data)
	//	} else {
	//		break
	//	}
	//}
}
```

# Golang中channel有缓冲与无缓冲的区别

## 无缓冲

1、定义时若未指定缓冲大小或设置为0，表示当前chan无缓冲。
2、在向chan写入数据时，会阻塞当前协程，直到其他协程从该chan中读取了数据。
3、基于定义2规则，无缓冲chan不能在一个协程同时进行读取与写入操作。
4、应用场景：主任务需要等待子任务的执行结果，可以使用无缓冲chan，主任务会阻塞等待子任务的执行结果写入chan后再继续执行当前业务逻辑。
5、chan可以进行多次的读写操作，可以在一个协程中生产多个数据，另一个协程中消费多个数据。

```go
func main() {
	c1 := make(chan string)
	go func() {
		fmt.Println("the will send data to chan")
		time.Sleep(time.Second * 1)
		c1 <- "string"
	}()

	// c1是无缓冲chan，此次会阻塞等待c1写入数据
	time.Sleep(time.Second * 2)
	s1 := <-c1
	fmt.Println("receive data with ", s1)
}
```

## 有缓冲

1、定义时需指定缓冲大小且大于0.
2、向chan写入数据时，若chan未满不会阻塞协程，满时阻塞线程直至缓冲有空间可写入。
3、写chan读取数据亦然。

```go
unc main() {
	c1 := make(chan string, 2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("the will send data to chan")
			// 当缓冲区满时，此时写入会阻塞
			c1 <- "string1"
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		// 当缓冲区空时，此时读取会阻塞
		s1 := <-c1
		fmt.Println("receive data with ", s1)
	}
}
```

### golang 并发编程之 select

* select是 go 中的一个控制结构，类似于 switch 语句，用于处理异步 IO 操作，select 会监听 case 语句中 channel 的读写操作，当 case中 channel 读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。
* select 中的 case 语句必须是一个 channel 操作
* select 中的 default 子句总是可运行的
* 如果有多个 case 都可以运行，select 会随机公平的选出一个执行，其他不会执行
* 如果没有可运行的 case 语句，且有 default 语句，那么就会执行 default 的动作
* 如果没有可运行的 case 语句，且没有 default语句，select将阻塞，直到某个 case 通信可以运行
* 通道关闭后，会随机进入一个 case，读出来的会是默认值。通道没关闭，会进入 default，通道没关闭并且没写 default 会死锁

```go
package _select

import (
	"fmt"
	"time"
)

func LSelect() {
	cI := make(chan int)
	cS := make(chan string)

	go func() {
		cI <- 1
		cS <- "邓文杰"

		defer close(cI)
		defer close(cS)
	}()

	for {
		select {
		case cIData := <-cI:
			fmt.Printf("%v\n", cIData)
		case cSData := <-cS:
			fmt.Printf("%v\n", cSData)
		default:
			fmt.Printf("%v\n", "default")
		}
		time.Sleep(time.Second)
	}
}
```

