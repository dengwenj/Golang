package main

import "fmt"

func main() {
	/*
		go 语言的变量声明

		显式声明 var name string = "邓文杰"
		var 关键字 name 变量名 string 变量类型 "邓文杰" 变量值

		隐式声明 age := 22
		省略了变量 var 省略了变量类型 number 用 : 代替，系统会自动的去找
	*/
	var name string = "邓文杰"
	age := 22 // 用到的多一点

	fmt.Println(name)
	fmt.Println(age)
}
