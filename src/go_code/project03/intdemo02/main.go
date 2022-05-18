package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		int 有符合
		uint 无符合
		rune 有符合 相当于 int32
		byte 无符合 相当于 uint8
	*/
	fmt.Println()

	// 整形的使用细节
	var n1 = 100
	// 查看 n1 的数据类型
	// fmt.PrintF() 可以用于做格式化输出
	fmt.Printf("n1 的数据类型是 %T\n", n1)

	// 如何在程序中查看某个变量的占用字节大小和数据类型
	var n2 int8 = 22
	// unsafe.Sizeof(n2) 是 unsafe 包的一个函数，可以返回 n2 变量占用的字节数
	fmt.Printf("n2 的数据类型是 %T n2 占用的字节数是 %d\n", n2, unsafe.Sizeof(n2))

	// go 程序中整型变量在使用时，遵守保小不保大原则
	// 0 ~ 255
	var age byte = 17
	fmt.Println(age)
}
