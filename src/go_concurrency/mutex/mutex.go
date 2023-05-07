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
