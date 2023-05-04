package _interface

import "fmt"

type Animal interface {
	eat(string) string
	say()
}

type Dog struct {
	Name string
}

type Cat1 struct {
	Name string
}

type Person struct {
	Name string
}

func (d Dog) eat(e string) string {
	d.Name = "狗狗"
	println("我是狗狗 eat")
	return e
}
func (d Dog) say() {
	println("我是狗狗 say")
}

func (c Cat1) eat(string2 string) string {
	println("我是猫猫 eat")
	return string2
}

func (c Cat1) say() {
	println("我是猫猫 say")
}

func (p Person) pp(a Animal) {
	fmt.Printf("%v\n%T\n", a, a)
	a.eat("狗狗")
	a.say()
}

func LInterfaceOCP() {
	p := Person{}
	p.pp(new(Dog))
	p.pp(new(Cat1))
}
