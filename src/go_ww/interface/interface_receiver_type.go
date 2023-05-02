package _interface

import "fmt"

type Pet interface {
	eat()
	say()
}

type Cat struct {
	Name string
}

//func (c Cat) eat() {
//	c.Name = "tom"
//}
//
//func (c Cat) say() {
//	c.Name = "kk"
//}

func (c Cat) eat() {
	c.Name = "tom"
}

func (c *Cat) say() {
	c.Name = "kk"
}

func LInterfaceReceiverType() {
	c := Cat{
		Name: "mm",
	}
	//var p Pet = c
	//p.eat()
	//p.say()
	//
	//fmt.Printf("%v\n", c) // {mm}

	var p Pet = &c
	p.eat()
	p.say()

	fmt.Printf("%v\n", c) // {kk}
}
