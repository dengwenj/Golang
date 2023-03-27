package learnconst

import "fmt"

func LearnConst() {
	//const constantName [type] = value
	const number float64 = 31.11
	const str = "你好"
	fmt.Printf("number: %v\n", number)
	fmt.Printf("str: %v\n", str)

	//	const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		a1 = 100
		a2 = 200
		a3
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	//	iota
	// iota 比较特殊，可以被认为是一个可被编译器修改的常量，它默认开始值是 0，每调用一次加 1，
	// 遇到 const 关键字时被重置为 0 !!!!
	const a4 = iota // 0
	const a5 = iota // 0

	const (
		a6 = iota // 0
		a7 = iota // 1
		a8 = iota // 2
	)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
	fmt.Printf("a7: %v\n", a7)
	fmt.Printf("a8: %v\n", a8)

	// 使用 _ 跳过某些值
	const (
		a9 = iota // 0
		_
		a10 = iota // 2
	)
	fmt.Printf("a9: %v\n", a9)
	fmt.Printf("a10: %v\n", a10)

	//	iota 声明中间插队  和跳过的道理是一样的
	const (
		a11 = iota // 0
		a12 = 1200 // 1200
		a13 = iota // 2
	)
	fmt.Printf("a11: %v\n", a11)
	fmt.Printf("a12: %v\n", a12)
	fmt.Printf("a13: %v\n", a13)
}
