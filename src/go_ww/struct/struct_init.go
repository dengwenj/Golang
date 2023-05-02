package _struct

import "fmt"

type Person1 struct {
	id, age     int
	name, email string
	hobby       []string
}

func LStructInit() {
	//结构体初始化
	var p = Person1{
		id:    2023,
		age:   23,
		name:  "邓文杰",
		email: "coderdwj@163.com",
		hobby: []string{"写博客", "听赵雷的歌曲"},
	}

	fmt.Printf("%v\n", p) // {2023 23 邓文杰 coderdwj@163.com [写博客 听赵雷的歌曲]}

	//使用列表的方式进行初始化
	p1 := Person1{
		2023,
		23,
		"邓文杰",
		"coderdwj@163.com",
		[]string{"民谣"},
	}
	fmt.Printf("%v\n", p1) // {2023 23 邓文杰 coderdwj@163.com [民谣]}

	//	部分初始化
	p2 := Person1{
		name: "邓文杰",
		age:  23,
	}
	fmt.Printf("%v\n", p2) // {0 23 邓文杰  []}
}
