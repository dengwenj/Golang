package slice

import "fmt"

func SliceCRUD() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//添加
	s1 = append(s1, 10)
	fmt.Printf("%v\n", s1)

	//删除 比如删除 4 5 6
	s1 = append(s1[:3], s1[6:]...)
	fmt.Printf("%v\n", s1)

	//更新
	s1[0] = 100
	fmt.Printf("%v\n", s1)

	//查询 找出索引是 2 的
	var key = 2
	for i, _ := range s1 {
		if i == key {
			fmt.Printf("%v\n%v\n", i, "找到了")
			break
		}
	}

	// copy 切片 深拷贝
	s2 := make([]int, len(s1))
	fmt.Printf("%v\n", s2)
	copy(s2, s1)
	s2[0] = 5201
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}
