package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		string 类型转基本数据类型
		使用 strconv 包的函数
	*/
	var strToI = "100"
	var num int64
	var num1 int8

	// 转成整数
	num, _ = strconv.ParseInt(strToI, 10, 64)
	num1 = int8(num)
	fmt.Printf("%T %v\n", num1, num1)

	// 转成小数
	var strToF = "123.456"
	var f float64
	f, _ = strconv.ParseFloat(strToF, 64)
	fmt.Printf("%T %v\n", f, f)

	// 转成 bool
	var strToB = "true"
	var b bool
	b, _ = strconv.ParseBool(strToB)
	fmt.Printf("%T %v\n", b, b)
}
