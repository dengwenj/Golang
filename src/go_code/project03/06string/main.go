package main

import "fmt"

func main() {
	// 字符串就是一串固定长度的字符连接起来的字符序列
	// go 的字符串是由单个字节连接起来的
	// go 语言的字符串的字节使用 utf-8 编码标识 Unicode 文本
	str := "你好世界 hello world"
	fmt.Println(str)

	str1 := `
		package main

		import (
			"fmt"
			"unsafe"
		)
		
		func main() {
			b := true
			// bool 类型占用存储空间是 1 个字节
			// bool 类型只能去 true 和 false
		
			fmt.Println("b 的占用空间是", unsafe.Sizeof(b))
		}
	`
	fmt.Println(str1)

	str2 := "hello12332112345677887921" + "hello12332112345677887921" +
		"hello12332112345677887921" + "hello12332112345677887921" +
		"hello12332112345677887921"
	fmt.Println(str2)

	// 基本数据类型的默认值
	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false
	fmt.Println(a, b, c, d)
}
