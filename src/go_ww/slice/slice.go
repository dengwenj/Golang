package slice

import "fmt"

func Slice() {
	// 切片（slice）是一个拥有相同类型元素的可变长度的序列。
	// 声明切片
	var s1 []string
	var s2 = []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)

	//使用 make 函数声明切片
	s3 := make([]string, 2)
	s3[0] = "李红"
	s3[1] = "王波"
	fmt.Printf("%v\n", s3)

	s4 := make([]string, 3, 3)
	s4[1] = "哈哈"
	fmt.Printf("%v\n", s4) // [ "哈哈" ]
	//切片长度
	fmt.Printf("%v\n", len(s4)) // 3
	//切片容量
	fmt.Printf("%v\n", cap(s4)) // 3
}
