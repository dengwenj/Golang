package flow

import "fmt"

func IfFlow() {
	/**
	1、不需要使用括号将条件包含起来
	2、大括号 {} 必须存在，及时只有一行语句
	3、左括号必须在 if 或 else 的同一行
	4、在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ; 进行分隔
	*/

	//初始变量可以声明在布尔表达式里面，但是作用域就在 if 语句里面了
	if age := 17; age < 18 {
		fmt.Println("不能干某些事")
	} else {
		fmt.Println("可以干某些事")
	}

	//不能使用 0 或 非 0 表示真假
	//number := 1
	//if number {
	//
	//}
}
