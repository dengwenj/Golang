package _struct

import "fmt"

type Person struct {
	name, email string
	age         int
}

func LStruct() {
	var s Person

	fmt.Printf("%v\n", s) // {  0}

	s.name = "邓文杰"
	s.age = 23
	s.email = "coderdwj@163.com"

	fmt.Printf("%v\n", s) // {邓文杰 coderdwj@163.com 23}

	//	匿名结构体
	var dwj struct {
		age   int
		hobby []string
	}
	dwj.age = 18
	dwj.hobby = []string{"社会人", "睡觉"}
	fmt.Printf("%v\n", dwj) // {18 [社会人 睡觉]}
}
