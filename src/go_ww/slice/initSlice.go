package slice

import "fmt"

func InitSlice() {
	//	初始化切片
	s5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", s5)

	//使用数组初始化
	arr := [...]int{1, 2, 3, 4}
	s6 := arr[:]
	fmt.Printf("%v\n", s6)

	//截取
	s7 := []int{1, 2, 3, 4, 5, 6}
	s8 := s7[:3]
	s9 := s7[3:]
	s10 := s7[2:5]
	s11 := s7[:]
	fmt.Printf("%v\n", s8)
	fmt.Printf("%v\n", s9)
	fmt.Printf("%v\n", s10)
	fmt.Printf("%v\n", s11)
}
