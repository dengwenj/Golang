package main

import "fmt"

func main() {
	/*
		go 在不同类型的变量之间赋值时需要显示转换，也就是说 go 中数据类型不能自动转换
		基本语法
			表达式 T(v) 将值 v 转换为类型 T
			T: 就是数据类型，比如 int32 int64
			v: 就是需要转换的变量
	*/
	a := 100
	var b float64 = float64(a)
	fmt.Println(b)
	var c int8 = int8(a)
	fmt.Println(c)

	/*
		细节说明
			1 go 中，数据类型的转换可以是从表示范围小 -> 表示范围大，也可以范围大 -> 范围小
			2 被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
			3 在转化中，比如将 int64 转成 int8 【-128 ~ 127】编译时不会报错，只是转换的结果是按溢出处理，和希望的结果不一样
	*/
	var num int64 = 134
	var num1 int8
	num1 = int8(num)
	fmt.Println(num1) // 不是 134 因为溢出了

	var n1 int32 = 12
	var n2 int64
	var n3 int8
	// n2 = n1 + 10 // 类型不一样
	// n3 = n1 + 10 // 类型不一样
	n2 = int64(n1) + 10
	n3 = int8(n1) + 10
	fmt.Println(n2, n3)

	// var n4 int32 = 22
	// var n5 int8
	// var n6 int8
	// n5 = int8(n4) + 127 // 溢出了 127 + 22 大于了 127 结果不是 149
	// n6 = int8(n4) + 128 // 128 直接溢出了
}
