package function

import "fmt"

func ff1(name string, callback func(str string)) {
	callback(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

type fn func(int2 int, int3 int) int

func cal(o string) fn {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func LfuncG() {
	ff1("邓文杰", func(str string) {
		fmt.Printf("%v\n", str) // 邓文杰
	})

	n := cal("+")(2, 3)
	fmt.Printf("%v\n", n)

	n1 := cal("-")(2, 1)
	fmt.Printf("%v\n", n1)
}
