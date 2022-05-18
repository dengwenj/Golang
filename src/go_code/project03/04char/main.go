package main

import "fmt"

func main() {
	// 当要存储字符是选用 byte
	// Go 中没有专门的字符类型，如果要存储单个字符（字母），一般使用 byt 来保存
	var b1 byte = 'a'
	var b2 byte = '0'
	// 格式化的形式输出对应的字符
	fmt.Println(b1, b2)         // 97 48
	fmt.Printf("%c %c", b1, b2) // a 0

	var b3 int = '邓'
	fmt.Printf("%c %d\n", b3, b3) // 邓 37011 用 byte 就存储不下，可以用 int 存储

	/*
		使用细节
			1 字符常量是用单引号括起来的单个字符，var b byte = 'a' var d int = '邓'
			2 go 中允许使用转义字符 \ 来将其后的字符传变为特殊字符型常量
			3 go 语言的字符使用 UTF-8 编码 英文字母 1 个字节，汉字 3 个字节
			4 go 中，字符的本质就是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值
			5 可以直接给某个变量赋一个数字，然后按格式化输出时 %c,会输出该数字对应的 Unicode 字符
			6 字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码
	*/
	var b4 = 10 + '邓'
	fmt.Println(b4) // 37021
}
