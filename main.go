package main

import "fmt"

func main() {
	dwj := "dengwenjie"
	num := 1209

	if dwj == "dengwenjie" {
		fmt.Println("进入了这里")
	} else if dwj == "ww" {
		fmt.Println("ww")
	} else {
		fmt.Println("进到 else 了")
	}

	switch num {
	case 1:
		fmt.Println("到了1")
	case 1209:
		fmt.Println("到了2")
		fallthrough // 是想让他走到下一个去
	case 3:
		fmt.Println("到了3")
	default:
		fmt.Println("go 语言里面默认会 break 终止")
	}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//var a = 0 // a := 0
	//for a < 10 {
	//	a++
	//
	//	if a == 2 {
	//		//return
	//		continue // 跳出本次循环
	//	}
	//
	//	fmt.Println(a)
	//}

	a := 0
A:
	for a < 10 {
		a++

		if a == 2 {
			//return
			break A // 跳转 A
			goto B
		}

		fmt.Println(a)
	}
B:
	fmt.Println("跳转到 B 来了")
}
