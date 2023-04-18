package arr

import "fmt"

func Element() {
	//访问数组元素

	var arr1 = [3]int{1, 2, 3}
	println(arr1[0])
	println(arr1[1])
	println(arr1[2])
	//赋值
	arr1[0] = 100
	println(arr1[0])

	arr2 := [...]string{"李雷", "韩梅梅", "王小波", "李银河"}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	arr3 := [2]bool{true, false}
	for i, v := range arr3 {
		fmt.Printf("arr3[%v]: %v\n", i, v)
	}
}
