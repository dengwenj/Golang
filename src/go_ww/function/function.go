package function

import "fmt"

func f1(a int, b int) (c int) {
	c = a + b
	return c
}

func f2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LFunction() {
	var c = f1(1, 2)
	fmt.Printf("%v\n", c)

	var d = f2(1, 2)
	fmt.Printf("%v\n", d)
}
