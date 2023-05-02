package pointer

import "fmt"

func LPointerArr() {
	//表示数组里面的元素的类型是指针类型
	var p1 [3]*int
	arr := [3]int{1, 2, 3}

	fmt.Printf("%v\n", p1) // [<nil> <nil> <nil>]
	for i := 0; i < len(arr); i++ {
		p1[i] = &arr[i]
	}
	fmt.Printf("%v\n", p1) // [0x140000ac000 0x140000ac008 0x140000ac010]

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\n", *p1[i])
	}
}
