package main

import "fmt"

func test() bool {
	fmt.Println("test...")
	return true
}

func main() {
	var n1 int8 = 100
	var n2 int8 = 12

	//fmt.Println(n1 > n2)
	// == != < > <= >= 结果是布尔值

	if n1 < 90 && test() {
		fmt.Println("&&。。。")
	}
	if n1 > 100 || test() {
		fmt.Println("||...")
	}

	if n1 > 10 && n2 < 10 {
		fmt.Println("没有进来")
	}

	if n1 < 110 || n2 > 100 {
		fmt.Println("进来了")
	}

	if !(n1 > 10) {
		fmt.Println("没有进来")
	}
}
