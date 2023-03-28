package string2

import (
	"bytes"
	"fmt"
	"strings"
)

func LearnString() {
	//字符串拼接
	var str string = "你好呀，邓文杰"
	str1 := "你好呀，韩梅梅"
	fmt.Printf("%v\n", str+str1)
	str2 := `
		你好呀
		你好哦
		你好啊
	`
	fmt.Printf("%v\n", str2)
	msg := fmt.Sprintf("name=%v&str1=%v", str, str1)
	fmt.Printf("%v\n", msg)

	var buffer bytes.Buffer
	buffer.WriteString("邓文杰")
	buffer.WriteString("18")
	buffer.WriteString("男")
	fmt.Printf("%v\n", buffer.String())

	//	字符串转义字符
	// \r 回车符 \n换行符 \t制表符 \'单引号 \"双引号 \\反斜杠
	str3 := "你好\n啊"
	fmt.Printf("%v\n", str3)
	str4 := "c:\\dengwj\\ww"
	fmt.Printf("%v\n", str4)

	//	字符串的切片操作
	str5 := "nihao,dengwj,ww"
	fmt.Printf("%c\n", str5[2])   // h
	fmt.Printf("%v\n", str5[2:5]) // hao [) 前面包括，后面不包括
	fmt.Printf("%v\n", str5[:5])  // nihao
	fmt.Printf("%v\n", str5[5:])  // ,dengwj

	// 字符串方法
	//获取长度
	fmt.Printf("%v\n", len(str5)) // 15
	//转换成 数组
	fmt.Printf("%v\n", strings.Split(str5, ",")) // [nihao dengwj ww]
	//判断该字符串是否存在里面
	fmt.Printf("%v\n", strings.Contains(str5, "niww"))  // false
	fmt.Printf("%v\n", strings.Contains(str5, "nihao")) // true
	//前缀
	fmt.Printf("%v\n", strings.HasPrefix(str5, "nihao")) // true
	//后缀
	fmt.Printf("%v\n", strings.HasSuffix(str5, "dengwj")) // false
	//获取该字符串在当中的索引
	fmt.Printf("%v\n", strings.Index(str5, "dengwj")) // 索引为 6
	fmt.Printf("%v\n", strings.LastIndex(str5, "w"))  // 索引为14
}
