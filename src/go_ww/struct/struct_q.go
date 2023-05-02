package _struct

import "fmt"

func LStructQ() {
	//结构体嵌套
	type Dog struct {
		name  string
		color string
	}

	type Person struct {
		name string
		age  int
		dog  Dog
	}

	p := Person{
		name: "莫言",
		age:  18,
		dog: Dog{
			"小花",
			"红色",
		},
	}
	fmt.Printf("%v\n", p) // {莫言 18 {小花 红色}}
}
