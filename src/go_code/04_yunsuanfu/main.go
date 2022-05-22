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

	a := 10
	b := 20
	fmt.Println(a, b)

	t := a
	a = b
	b = t
	fmt.Println(a, b)

	c := 30
	d := 40

	c = c + d
	d = c - d
	c = c - d
	fmt.Println(c, d)
	//c, d = d, c
	//fmt.Println(c, d)
}
