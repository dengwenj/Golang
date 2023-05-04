package _interface

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
	Name string
}

func (f Fish) fly() {
	println("我是飞")
}

func (f Fish) swim() {
	println("我是鱼")
	f.Name = "我啥谁"
}

func LInterfaceCompose() {
	// 接口类型
	var f FlyFish
	f = &Fish{}
	f.swim()
	f.fly()
}
