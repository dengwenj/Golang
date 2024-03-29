package _interface

// USB 定义接口
type USB interface {
	read()
	write()
}

// Computer M 它们用了共同的方法
type Computer struct {
	Name string
}

type M struct {
	Name string
}

// 方法
func (m M) read() {
	println("我是m read")
}

func (m M) write() {
	println("我是m read")
}

func (c Computer) read() {
	println("我是c read")
}

func (c Computer) write() {
	println("我是c write")
}

func LInterface() {
	c := Computer{
		Name: "电脑",
	}
	//var u USB = c
	// 用的时候以接口为类型，用结构体初始化
	var u USB = Computer{
		Name: "电脑",
	}
	u.read()
	u.write()

	m := M{
		Name: "手机",
	}

	test(c)
	test(m)
}

func test(u USB) {
	u.read()
	u.write()
}
