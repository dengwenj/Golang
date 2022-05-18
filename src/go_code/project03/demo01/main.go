package main

import "fmt"

// 全局的变量
var uName = "dengwenjie"
var age = 22
var sex = "男"

func main() {
	/*
		变量表示内存中的一个存储区域
		该区域有自己的名称（变量名）和类型（数据类型）
		变量如果没有赋初始值，会有默认值，int 默认值为 0 string 默认值为空串。小数默认值为 0
		同一个作用域下不能重名
	*/
	var hh string
	var num int = 1
	// 一次声明多个变量的方式
	var num1, num2, num3 = 11, 22, 33
	// := 方式声明
	name := "dengwj"
	hh = "ww"
	// 重新赋值
	name = "dwj"

	//也可以
	var (
		name1 = uName
		age1  = age
		sex1  = sex
	)

	var str1 = "dwj"
	var str2 = "ww"
	var n1 = 1
	var n2 = 2
	fmt.Println(str1 + str2)
	fmt.Println(n1 + n2)

	fmt.Println(num, num1, num2, num3, name, hh)
	fmt.Println(name1, age1, sex1)
}
