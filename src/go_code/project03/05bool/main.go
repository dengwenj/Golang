package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := true
	// bool 类型占用存储空间是 1 个字节
	// bool 类型只能去 true 和 false

	fmt.Println("b 的占用空间是", unsafe.Sizeof(b))
}
