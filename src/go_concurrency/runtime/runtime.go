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
