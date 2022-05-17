package main

import "fmt"

func main() {
	/*
		转义
		\t 制表
		\n 换行符
		\\ 一个\
		\" 一个"
		\r fmt.Println("王小波和李银河\r韩梅梅") // 韩梅梅和李银河
	*/
	fmt.Println("dwj\tww") // dwj     ww
	fmt.Println("你好\n世界")  // 你好
	//                               世界
	fmt.Println("c:\\wjd\\dd")      // c:\wjd\dd
	fmt.Println("w说\"i love you\"") // w说"i love you"
	fmt.Println("王小波和李银河\r韩梅梅")     // 韩梅梅和李银河

	fmt.Println("姓名\t年龄\t籍贯\t住址\ndwj\t22\t重庆\t上海")
}
