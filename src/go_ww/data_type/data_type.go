package learndatatype

import "fmt"

func fn() {

}

func LearnDataType() {
	var name string = "邓文和"
	age := 18
	b := true
	// %T 是打印类型 \n 是换行
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)

	//	指针类型
	var a = 1
	dwj := &a
	fmt.Printf("%T\n", dwj)

	//	数组类型
	arr := [2]int{1, 2}
	fmt.Printf("%T\n", arr)

	// 切片 类型, 就是不写长度
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", s)

	// 函数类型
	fmt.Printf("%T\n", fn)
}
