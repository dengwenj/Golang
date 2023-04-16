package arr

import "fmt"

func Array() {
	var arr1 [2]int
	var arr2 [2]string
	var arr3 [2]bool
	arr1[1] = 2
	arr2[1] = "邓文杰"
	arr3[1] = true
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr4 := [3]int{1, 2, 3}
	fmt.Println(arr4)

	arr5 := [...]string{"李雷", "韩梅梅", "马伊琍"}
	fmt.Println(arr5)

	arr6 := [...]int{1: 11, 3: 133}
	fmt.Println(arr6)
}
