package format

import "fmt"

// Website 结构体
type Website struct {
	Name string
	Age  int
}

func Format() {
	a := Website{
		Name: "你好啊",
		Age:  12,
	}

	b := true

	number := 123

	fmt.Printf("%v\n", a)      // {你好啊 12}
	fmt.Printf("%#v\n", a)     // format.Website{Name:"你好啊", Age:12}
	fmt.Printf("%T\n", number) // int 类型
	fmt.Printf("%t\n", b)      // true
}
