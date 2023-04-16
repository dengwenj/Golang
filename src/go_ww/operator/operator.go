package operator

import "fmt"

func Operator() {
	a := 100
	b := 200

	c := a + b
	c = a - b
	c = a * b
	c = a / b
	fmt.Printf("%v\n", c)

	/*
	 + - * /
	 == !=
	 && || !
	 =
	 & | << >>
	*/
}
