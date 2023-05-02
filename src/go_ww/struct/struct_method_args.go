package _struct

import "fmt"

type Person4 struct {
	name string
}

// 值类型
func (p Person4) who() {
	p.name = "不会修改"
}

// 指针类型
func (p *Person4) who2() {
	// p.name 自动解引用
	// (*p).name = "会修改"
	p.name = "会修改"
}

func test(p1 Person4) {
	p1.name = "不会修改"
}

func test3(p1 *Person4) {
	p1.name = "会修改哦"
}

func LStructMethodArgs() {
	p := Person4{
		"邓文杰",
	}
	p.who()
	fmt.Printf("%v\n", p) // {邓文杰}
	p.who2()
	fmt.Printf("%v\n", p) // {会修改}

	p1 := Person4{
		"莫言",
	}
	test(p1)
	fmt.Printf("%v\n", p1) // {莫言}

	test3(&p1)
	fmt.Printf("%v\n", p1) // {会修改哦}
}
