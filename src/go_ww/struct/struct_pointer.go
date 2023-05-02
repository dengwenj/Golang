package _struct

import "fmt"

func LStructPointer() {
	p := Person1{
		name: "Deng",
		age:  23,
	}
	//结构体指针
	var pP *Person1
	pP = &p
	fmt.Printf("%v\n", pP)  // &{0 23 Deng  []}
	fmt.Printf("%p\n", pP)  // 0x14000100050
	fmt.Printf("%v\n", *pP) // {0 23 Deng  []}

	//使用 new 关键字创建结构体指针
	var pP1 = new(Person1)
	pP1 = &p
	//(*pP1).id = 101
	pP1.id = 101
	pP1.hobby = []string{"吃饭"}
	fmt.Printf("%p\n", pP1)  // 0x1400008c050
	fmt.Printf("%v\n", *pP1) // {101 23 Deng  [吃饭]}
}
