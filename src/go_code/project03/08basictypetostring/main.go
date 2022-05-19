package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		基本类型转 string 类型
		方式1：fmt.Sprintf("%参数", 表达式)
		fmt.Sprintf() 会返回转换后的字符串
	*/
	var num1 int32 = 100
	var num2 float64 = 12.09
	var b bool = true
	var str string
	// fmt.Sprintf 方式
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("%T %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("%T %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("%T %q\n", str, str)

	// strconv 方式
	var num3 int = 200
	var num4 float64 = 123.456
	var b2 bool = false

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("%T %q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("%T %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("%T %q\n", str, str)

	// int 转换为 string
	var num5 = 300
	str = strconv.Itoa(num5)
	fmt.Printf("%T %q\n", str, str)
}
