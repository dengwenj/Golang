package _struct

import "fmt"

type Person3 struct {
	name string
	say  func()
}

// 结构体方法，struct 的接收者
func (p Person3) eat() {
	fmt.Printf("%v\n", p.name)
	fmt.Printf("%v\n", "我是Person3 结构体的方法")
}

func (p Person3) login(username string, password string) bool {
	if username == "邓文杰" && password == "123456" {
		fmt.Printf("%v\n", p.name)
		return true
	}
	return false
}

func LStructMethod() {
	p1 := Person3{
		name: "邓文杰",
		say: func() {
			// 这样写拿不到属性，?目前不知道
			println("进")
		},
	}
	p1.eat()
	p1.say()

	isT := p1.login("邓文杰", "123456")
	if isT {
		println("登录成功")
	} else {
		println("用户名或密码错误")
	}
}
