package learnvar

import "fmt"

// (name3 string, name4 string) 不写，可以是匿名变量
func test() (string, string) {
	// 初始化多个变量
	var dwj, hmm, ll = "邓文杰", "韩梅梅", "李雷"
	fmt.Printf("name1: %v\n", dwj)
	fmt.Printf("name1: %v\n", hmm)
	fmt.Printf("name1: %v\n", ll)

	//批量声明,使用一个 var 关键字，把一些变量写在一个括号()里
	var (
		hh string
		xx = "嘻嘻"
	)
	fmt.Printf("name1: %v\n", hh)
	fmt.Printf("name1: %v\n", xx)

	// 短变量声明 只能在函数内部使用，函数外部不能使用！！！！
	// 在函数内部，可以使用 := 运算符对变量进行声明和初始化
	study := "学习"
	gg := "GG"
	fmt.Printf("name1: %v\n", study)
	fmt.Printf("name1: %v\n", gg)

	// 匿名变量
	// 如果我们接收多个变量，有一些变量使用不到，可以使用下划线_表示变量名称，这种变量叫做匿名变量
	return "王小波", "李银河"
}

func LearnVar() {
	// 变量初始化语法
	// var 变量名 类型 = 表达式
	var name string = "dengwenjie"
	var age int = 23
	var gender string = "男"
	fmt.Printf("name1: %v\n", name)
	fmt.Printf("name1: %v\n", age)
	fmt.Printf("name1: %v\n", gender)

	name1, _ := test()
	fmt.Printf("name1: %v\n", name1)
}
