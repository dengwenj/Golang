package main

import (
	"Golang/testpackage"
	"fmt"
)

func main1() {
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
	fmt.Println(testpackage.Test1)
	fmt.Println(testpackage.Test2)
}

/*
	包：
	一个文件夹下面只能有一个包，但是可以有多个文件，每个文件都是这个包，只不过分开写了
	他们是属于一个包的，所以变量名是不可以重复的
	一个包可以把所有文件的内容整合到一起，只是视觉上看起来是多个文件
*/
