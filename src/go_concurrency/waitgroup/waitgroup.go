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
