package _type

import "fmt"

func LTypeDefine() {
	//类型定义
	type NewInt int
	var num NewInt = 100

	fmt.Printf("%v%T\n", num, num) // 100 _type.NewInt

	//	类型别名
	type ww = string
	var str ww
	str = "ww"
	fmt.Printf("%v%T\n", str)
}
