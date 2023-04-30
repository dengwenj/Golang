package _map

import "fmt"

func LMap() {
	//var m map[string]string
	// 需要通过make方法分配确定的内存地址
	m := make(map[string]string)
	fmt.Printf("%v\n", m) // map[]
	m["name"] = "邓文杰"
	fmt.Printf("%v\n", m) // map[name:邓文杰]

	m2 := map[string]string{
		"name":   "文杰",
		"age":    "23",
		"gender": "男",
	}
	fmt.Printf("%v\n", m2)
	m2["hobby"] = "睡觉"
	fmt.Printf("%v\n", m2)

	//判断某个键是否存在
	v, ok := m2["name"]
	fmt.Printf("%v\n", v)  // 文杰
	fmt.Printf("%v\n", ok) // true

	v, ok = m2["ww"]
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", ok) // false

	//	遍历
	for i, v := range m2 {
		fmt.Printf("%v\n%v\n", i, v)
	}
}
