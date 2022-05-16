package main

import (
	"fmt"
	"strconv"
)

func main2() {
	var num1 uint
	var num2 int
	var num3 float64
	var str string
	var bool1 bool
	var bool2 bool

	num1 = 999
	num2 = -999
	num3 = 3.14
	str = "我是字符串"
	bool1 = false
	bool2 = true

	fmt.Println(num1, num2, num3, str, bool1, bool2)

	// fmt.Printf("%T", str, "zi") // 用于判断是什么类型

	// 字符串转成整数类型  string -> int
	str1 := "123"
	var int123, _ = strconv.Atoi(str1)
	fmt.Println(int123)
	fmt.Printf("%T", int123)

	// string -> int64
	str2 := "234"
	int2, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Println(int2)
	fmt.Printf("%T", int2)

	// int -> string
	num111 := 111
	string1 := strconv.Itoa(num111)
	fmt.Println(string1)
	fmt.Printf("%T", string1)

	// 字符串到 float32、float64
	str3 := "3.14"
	float3, _ := strconv.ParseFloat(str3, 32)
	fmt.Println(float3)

	// int64 转 int
	// int(64的)
}

/*
	基本数据类型

	int
		uint 正整数
		byte 需要字节存储时候才会用到 一般是字节组成的数组
		rune 等价于 int32 存储一个 Unicode 编码
		int 全部整数
	float 浮点型
	string 字符串
	bool 布尔类型

	复杂数据类型

		结构 struct
		接口 interface
		数组 [value, value]
    切片 slice
		map {key: value}
		指针
		函数 func
		管道 chan
*/
