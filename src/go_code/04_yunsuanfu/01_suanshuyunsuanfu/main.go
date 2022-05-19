package main

import "fmt"

func main() {
	// 如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4) // 2
	var n1 float32 = 10 / 4
	fmt.Println(n1) // 2

	// 如果希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10 / 4.0
	fmt.Println(n2) // 2.5

	// a % b = a - a / b * b
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(-10 % -3) // -1
	fmt.Println(10 % -3)  // 1

	// 在 go 中， ++ 和 -- 只能独立使用
	//var i = 10
	//var x int
	//x = i++ // 只能独立使用
	//i++
	//if i++ > 0 { // 只能独立使用
	//
	//}
}
