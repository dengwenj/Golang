package main

import "fmt"

func main() {
	/*
		基本数据类型，变量存的就是值，也叫值类型
		获取变量的地址，用 & 比如：var num int 获取 num 的地址：&num
		指针类型，变量存的是一个地址，这个地址指向的空间存的才是值，比如：var ptr *int = &num
		获取指针类型所指向的值，使用 *, 比如：var ptr *int, 使用 *ptr 获取 ptr 存储的内容(地址) 指向的值
		& 取地址 * 取值  *ptr 可以访问到那个空间
	*/
	var num = 1

	// ptr 是个指针变量
	var ptr *int = &num
	fmt.Println(ptr) // 0x1400001c070

	fmt.Println(*ptr) // 1

	var str = "ffff"
	var s = &str
	fmt.Println(s) // 0x14000104210

	var num1 int8 = 100
	var ptr1 *int8 = &num1
	fmt.Println(ptr1, num1) // 0x1400001c078 100
	// 修改 num 的值
	*ptr1 = 110
	fmt.Println(num1) // 110
}
