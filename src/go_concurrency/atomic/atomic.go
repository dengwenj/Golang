package atomic

import (
	"sync/atomic"
	"time"
)

//var lock sync.Mutex
var num int32 = 100 // 原子变量，这个变量是个原子的

func add() {
	//原子变量的方式
	// atomic 原子操作
	atomic.AddInt32(&num, 1) // 对这变量进行操作的时候保证他的线程安全
}

func sub() {
	atomic.AddInt32(&num, -1)
}

//func add() {
//	//用互斥锁的方式
//	lock.Lock() // 锁起来,其他协程运行不了
//	num++
//	lock.Unlock()
//}
//
//func sub() {
//	lock.Lock()
//	num--
//	lock.Unlock()
//}

func LAtomic() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)
	println(num)
}
