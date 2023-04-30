package function

import "fmt"

func f3() {
	fmt.Printf("%v\n", "f1 没有返回值")
}

func f4() string {
	return "有一个返回值"
}

func f5() (string, int) {
	return "多个返回值", 2
}

func f6() (name string, age int) {
	name = "邓文杰"
	age = 23
	return name, age
}

func f7() (name string, age int) {
	name = "余华"
	age = 18
	return // 等价于 return name, age
}

func LFuncReturn() {
	f3()
	str := f4()
	fmt.Printf("%v\n", str)
	str1, n := f5()
	fmt.Printf("%v\n%v\n", str1, n)
	name, age := f6()
	fmt.Printf("%v\n%v\n", name, age)
	name1, age1 := f7()
	fmt.Printf("%v\n%v\n", name1, age1)
}
