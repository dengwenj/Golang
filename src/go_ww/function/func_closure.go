package function

import "fmt"

func f20() func(int) int {
	var x int
	//用到了闭包 -> 闭包 = 函数 + 引用环境
	return func(num int) int {
		x += num
		return x
	}
}

type fn1 func(int) int

func cal1(base int) (fn1, fn1) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(b int) int {
		base -= b
		return base
	}
	return add, sub
}

func LFuncClosure() {
	//numFn := f20()
	//n := numFn(10)
	//fmt.Printf("%v\n", n) // 10
	//n = numFn(20)
	//fmt.Printf("%v\n", n) // 30
	//n = numFn(30)
	//fmt.Printf("%v\n", n) // 60
	//
	//numFn1 := f20()
	//n2 := numFn1(100)
	//fmt.Printf("%v\n", n2) // 100
	//n2 = numFn1(200)
	//fmt.Printf("%v\n", n2) // 300

	add, sub := cal1(100)
	res := add(1)
	fmt.Printf("%v\n", res) // 101
	res = sub(2)
	fmt.Printf("%v\n", res) // 99
	res = add(3)
	fmt.Printf("%v\n", res) // 102
	res = sub(4)
	fmt.Printf("%v\n", res) // 98
}
