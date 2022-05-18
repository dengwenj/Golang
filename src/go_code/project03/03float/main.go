package main

import "fmt"

func main() {
	num := 124.11111
	fmt.Printf("num 的数据类型数 %T", num) // float64 默认

	// 有固定的的范围和字段长度，不受具体 os 的影响
	// 科学计数法形式：5.1234e2 = 512.34 即 5.12 * 10的二次方
	// 通常情况下，应该使用 float64
}
