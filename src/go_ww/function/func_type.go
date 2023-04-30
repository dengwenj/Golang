package function

import "fmt"

type fun func(int, int) int

func f14(a int, b int) int {
	return a + b
}

func f16(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LFuncType() {
	var f15 fun
	f15 = f14
	num := f15(1, 2)
	fmt.Printf("%v\n", num)

	var f17 fun
	f17 = f16
	max := f17(1, 2)
	fmt.Printf("%v\n", max)
}
