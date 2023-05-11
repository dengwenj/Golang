### golang 标准库 os 模块-文件目录相关

* os 标准库实现了平台（操作系统）无关的编程接口

```go
package _os

func LFileDir() {
	//创建文件
	//file, err := os.Create("dengWJ.txt")
	//if err == nil {
	//	println(file)
	//}

	//	创建目录 创建单个目录
	//err := os.Mkdir("dwj", os.ModePerm)
	//if err != nil {
	//
	//}
	// 创建多个目录
	//os.MkdirAll("dwj/ww", os.ModePerm)

	//删除文件或目录
	//os.Remove("dengWJ.txt")
	//os.Remove("dwj")
	//os.RemoveAll("dwj")

	//	获取工作目录
	//dir, err := os.Getwd()
	//if err == nil {
	//	fmt.Printf("dir: %v\n", dir) // /Users/dengwenjie/Golang/Golang/src/go_std
	//}
	//修改工作目录
	//os.Chdir("/Users/dengwenjie/Golang/Golang")
	//dir, _ = os.Getwd()
	//fmt.Printf("new: %v\n", dir)

	//	获取临时目录
	//s := os.TempDir()
	//fmt.Printf("TempDir: %v\n", s)

	//	重命名文件
	//os.Rename("dWJ.txt", "dWj.txt")

	//	读取文件
	//file, _ := os.ReadFile("dWj.txt")
	//fmt.Printf("%v\n", string(file))

	//	写入文件
	//os.WriteFile("dWj.txt", []byte("你好啊"), os.ModePerm)
}
```

