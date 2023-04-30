package function

import "fmt"

func f21(n int) int {
	if n == 1 {
		return 1
	}
	// 函数调用栈，后进先出，上面的函数执行完先销毁
	return n * f21(n-1)
}

func LFuncRuc() {
	i := f21(5)
	fmt.Printf("%v\n", i) // 120
}
