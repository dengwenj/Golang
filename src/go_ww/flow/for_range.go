package flow

import "fmt"

//数组
func forRangeFlow1() {
	arr := [...]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("%v:%v\n", i, v)
	}
}

//切片
func forRangeFlow2() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range s {
		fmt.Printf("%v:%v\n", i, v)
	}
}

//map
func forRangeFlow3() {
	m := make(map[string]string, 0)
	m["name"] = "邓文杰"
	m["age"] = "23"
	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
}

//字符串
func forRangeFlow4() {
	str := "hello world"
	for i, v := range str {
		fmt.Printf("%v:%c\n", i, v)
	}
}

func ForRangeFlow() {
	forRangeFlow1()
	forRangeFlow2()
	forRangeFlow3()
	forRangeFlow4()
}
