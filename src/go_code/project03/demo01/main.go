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
	*/
	var num int = 1
	// 一次声明多个变量的方式
	var num1, num2, num3 = 11, 22, 33
	// := 方式声明
	name := "dengwj"

	//也可以
	var (
		name1 = uName
		age1  = age
		sex1  = sex
	)

	fmt.Println(num, num1, num2, num3, name)
	fmt.Println(name1, age1, sex1)
}
