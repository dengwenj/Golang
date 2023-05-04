package _interface

import "fmt"

type An struct {
	name string
	age  int
}

func (a An) eat(string2 string) {
	println(string2 + "实现继承")
}
func (a An) sleep(string2 string) {
	println(string2 + "实现继承 sleep")
}

type Cat3 struct {
	color string
	An    // 实现了继承
}

type Dog3 struct {
	ddd string
	An
}

// LStructExtends 实现继承
func LStructExtends() {
	c := Cat3{
		"白色",
		An{name: "小猫", age: 3},
	}
	fmt.Printf("%v\n", c.color)
	c.eat("猫猫")
	c.sleep("猫猫")

	d := Dog3{
		"ddd",
		An{
			name: "小狗",
			age:  4,
		},
	}
	fmt.Printf("%v\n", d.ddd)
	d.eat("狗狗")
	d.sleep("狗狗")
}
