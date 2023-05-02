package _struct

import "fmt"

type Person2 struct {
	name string
	age  int
}

//1、直接传递结构体
func test1(p Person2) {
	p.age = 101
	p.name = "韩梅梅"
	fmt.Printf("%v\n", p) // {韩梅梅 101}
}

//2、传递的是结构体指针 会修改原来的
func test2(pP *Person2) {
	pP.name = "李银河"
	pP.age = 11

	fmt.Printf("%v\n", *pP) // {李银河 11}
}

func LStructParams() {
	//结构体作为参数传递给函数
	p := Person2{
		name: "李雷",
		age:  18,
	}
	test1(p)
	fmt.Printf("%v\n", p) // {李雷 18}，不会改变，因为是值类型

	p1 := Person2{
		name: "王小波",
		age:  10,
	}
	pP := &p1
	test2(pP)              // test2(&p1)
	fmt.Printf("%v\n", p1) // {李银河 11}
}
