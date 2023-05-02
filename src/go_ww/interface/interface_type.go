package _interface

import "fmt"

// Pet1 一个类型可以实现多个接口
type Pet1 interface {
	p1()
}
type Pet2 interface {
	p2()
}

// I1 一个接口可以被多个类型实现
type I1 interface {
	say()
}

type Ppp struct {
}

type S1 struct {
}
type S2 struct {
}

func (p Ppp) p1() {
	fmt.Printf("%v\n", "11")
}

func (p Ppp) p2() {
	fmt.Printf("%v\n", "22")
}

func (s S1) say() {
	println("斤1")
}
func (s S2) say() {
	println("进2")
}

func LInterfaceType() {
	//var pp Pet1 = Ppp{}
	//var ppp Pet2 = Ppp{}
	//pp.p1()
	//ppp.p2()

	var i I1
	i = S1{}
	i.say()
	i = S2{}
	i.say()
}
