package _interface

import "fmt"

type Person4 struct {
	Name string
	Age  int
}

// NewPerson 模拟构造函数
func NewPerson(name string, age int) (*Person4, error) {
	if name == "" {
		return nil, fmt.Errorf("姓名不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("年龄不能小于0")
	}
	return &Person4{Name: name, Age: age}, nil
}

func LConstruct() {
	person, err := NewPerson("dengWj", 18)
	if err == nil {
		fmt.Printf("%v\n", person)
	} else {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", person.Name)
	fmt.Printf("%v\n", person.Age)
}
