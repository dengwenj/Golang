package main

import "fmt"

func main4() {
	// 把同一类元素放在一起的集合 数组
	// [数组的长度]元素类型{元素1, 元素2, ...}
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 不确定的情况下可以用 ...
	arr1 := [...]string{"dengwj", "ww", "hmm"}
	fmt.Println(arr1)

	// 循环数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 可以把索引和值遍历处理
	for idx, value := range arr {
		fmt.Println(idx, value)
	}

	// 定义数组先不赋值
	var arr3 [3]int
	fmt.Println(arr3)
	arr3 = [...]int{1, 3, 3}
	fmt.Println(arr3)

	var arr4 = new([10]int)
	fmt.Println(arr4)

	arr5 := [2]int{}
	fmt.Println(arr5)

	// 多维数组
	er := [2][3]int{
		{1, 2}, // 没有达到数组的长度就是用 0 代替
		{3, 4, 5},
	}
	fmt.Println(er)
}
