package function

import "fmt"

func f8(a int) {
	a = 200
}

func f9(args ...int) {
	fmt.Printf("%v\n%T", args, args)
}

func f10(name string, age int, args ...string) {
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", age)

	for i, v := range args {
		fmt.Printf("%v\n%v\n", i, v)
	}
}

func f11(s1 []int) {
	s1[0] = 20000
}

func LFuncParams() {
	i1 := 100
	f8(i1)
	fmt.Printf("%v\n", i1)

	f9(1, 2, 3, 4, 5)

	f10("邓文杰", 23, "韩梅梅", "李雷")

	var s = []int{1, 2, 3, 4, 5}
	f11(s)
	fmt.Printf("%v\n", s)
}
