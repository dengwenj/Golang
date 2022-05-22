package main

import "fmt"

func test() bool {
	fmt.Println("test...")
	return true
}

func main() {
	var n1 int8 = 100
	var n2 int8 = 12

	//fmt.Println(n1 > n2)
	// == != < > <= >= 结果是布尔值

	if n1 < 90 && test() {
		fmt.Println("&&。。。")
	}
	if n1 > 100 || test() {
		fmt.Println("||...")
	}

	if n1 > 10 && n2 < 10 {
		fmt.Println("没有进来")
	}

	if n1 < 110 || n2 > 100 {
		fmt.Println("进来了")
	}

	if !(n1 > 10) {
		fmt.Println("没有进来")
	}

	a := 10
	b := 20
	fmt.Println(a, b)

	t := a
	a = b
	b = t
	fmt.Println(a, b)

	c := 30
	d := 40

	c = c + d
	d = c - d
	c = c - d
	fmt.Println(c, d)
	//c, d = d, c
	//fmt.Println(c, d)

	var name string
	var age byte
	var sal float32
	var isPass bool

	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("是否通过")
	//fmt.Scanln(&isPass)

	//fmt.Printf("姓名%v\n 年龄%v\n 薪水%v\n 是否通过%v\n", name, age, sal, isPass)
	fmt.Println("请输入姓名、年龄、薪水、是否通过")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名%v\n 年龄%v\n 薪水%v\n 是否通过%v\n", name, age, sal, isPass)
}
