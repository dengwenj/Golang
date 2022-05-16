package main

import "fmt"

func main() {
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

	fmt.Printf("%T", str, "zi") // 用于判断是什么类型
}

/*
	基本数据类型

	uint 正整数
	int 整数
	float 浮点型
	string 字符串
	bool 布尔类型
*/
