package _interface

type Person2 struct {
	// 大写外面包可以访问，小写不可以
	name string
	age  int
}

func (p Person2) eat() {
	println("我可以吃")
}
func (p Person2) work() {
	println("我可以工作")
}

func LStructOOP() {
	p := Person2{
		name: "dengWJ",
		age:  18,
	}
	p.eat()
	p.work()
}
