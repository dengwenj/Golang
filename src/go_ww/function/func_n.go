package function

import "fmt"

func LFuncN() {
	//匿名函数
	f := func() {
		fmt.Printf("%v\n", "我是匿名函数")
	}
	f()

	//立即执行函数、自己调用自己
	(func() {
		fmt.Printf("%v\n", "我是立即执行函数")
	})()
}
