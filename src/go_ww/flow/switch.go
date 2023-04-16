package flow

func fn1() {
	var num = 1
	switch num {
	case 1, 2, 3, 4, 5:
		println("工作日")
	case 6, 7:
		println("休息日")
	}
}

func fn2() {
	num1 := 2
	switch {
	case num1 > 1:
		println("大于三")
	case num1 < 1:
		println("小于一")
	}
}

func fn3() {
	num2 := 2
	switch {
	case num2 > 1:
		println("大于三")
		fallthrough
	case num2 < 1:
		println("小于一")
	}
}

func SwitchFlow() {
	/**
	1、支持多条件匹配
	2、不同的 case 之间不使用 break 分隔，默认只会执行一个 case
	3、若想要执行多个 case，需要使用 fallthrough 关键字，也可以 break 终止
	4、分支还可以使用表达式，例如：a > 10
	*/

	fn1()
	fn2()
	fn3()
}
