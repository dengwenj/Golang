package pointer

import "fmt"

func LPointer() {
	//指针变量
	var p *int
	fmt.Printf("%v\n", p) // nil
	num := 100
	//拿到 num 的地址赋值给 p
	p = &num
	fmt.Printf("%v\n", p)  // 0x1400012c010
	fmt.Printf("%T\n", p)  // *int
	fmt.Printf("%v\n", *p) // 100
}
