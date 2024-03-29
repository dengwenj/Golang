



### Go 项目管理工具

**go应用使用包和模块来组织代码，包对应到文件系统就是文件夹，模块就是.go的go源文件。一个包中会有多个模块，或者多个子包**

#### 实现步骤

1、创建项目

2、初始化项目

3、创建包

4、创建模块

5、相互调用 

* 初始化项目 -> go mod init 模块名称(一般项目名称) 
* 文件夹相当于是个包，文件相当于是个模块 .go
* 入口程序：主包package main，主函数func main() {}

运行go中本地包的时候如果出现这个报错，网上很多说是需要修改GO111MODULE=on， 但是对于我来说都没用， 最后发现是因为命名导致的。 记录如下。

为了表示各个文件名和路径之间的无关联性，我尽量把命名随意化和多样化。

首先在桌面上建立一个文件夹aaa

在文件夹里运行下面的命令生成mod文件

```
go` `mod init aaa
```

**注意：这里的aaa和文件夹的名字aaa必须要一致。**

然后在aaa根目录里创建文件bbb.go作为主程序入口文件，里面的包名必须是main。

然后在aaa根目录创建子目录这里我命名为ccc，里面创建一个文件叫eee.go，包名可以叫ddd。 在vscode里展示如图下所示。

![img](https://img.jbzj.com/file_images/article/202303/202303170935536.png)

![img](https://img.jbzj.com/file_images/article/202303/202303170935537.png)

然后在aaa的根目录打开命令窗口， 运行go run bbb.go就会成功调用到子包里面的方法。

为了不报错，这里有几个重点

**第一，根目录文件夹名称xxx和go mod init xxx 必须要一致。**

**第二，主文件中import的是子包的目录路径，不能写子包的文件名或者包名。**

**第三，调用子包的方法的时候， 前缀必须是子包的包名（package名），和路径或者文件名无关。**

### go语言命名规范

* 命名规则设计变量、常量、全局变量、结构、接口、方法等的命名。go语言从语法层面进行了一下限定：任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。
* 当命名（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：GetUserName，那么使用这种形式的标识符的对象就可以被外部包的代码所使用，这被称为导出；命名如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的。

**包名称**

保持 `package`的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。包名应该为小写单词，不要使用下划线或者混合大小写。

```go
package dao
package service
```

**文件命名**

尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词。

```go
customer_dao.go
```

**结构体命名**

采用驼峰命名法，首字母根据访问控制大写或者小写

`struct` 申明和初始化格式采用多行，例如：

```go
type Customerorder struct {
	Name string
	Address string
}
order := CustomerOrder{"tom", "上海"}
```

**接口命名**

命名规则基本和上面的结构体类型

单个函数的结构名以“er”作为后缀，例如：`Reader`、`Writer`

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

**变量命名**

和结构体类型，变量名一般遵循驼峰法，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循一下规则：

如果变量为私有，且特有名词为首个单词，则使用小写，如appService 若变量为bool类型，则名称应以Has、Is、Can、Allow开头，

```go
var isExist bool
```

**常量命名**

常量均需使用全部大写字母组成，并使用下划线分词，

```go
const APP_URl = "https://dengwj.vip"
```

如果是枚举类型的常量，需要先创建相应类型

```go
type Scheme string
const (
	HTTP Scheme = "http"
	HTTPS Scheme = "https"
)
```

**错误处理**

错误处理的原则就是不能丢弃任何有返回 err 的调用，不要使用_丢弃，必须全部处理。接收到错误，要么返回 err，或者使用 log 记录下来尽早 return：一旦有错误发生，马上返回，尽量不要使用 panic，除非知道在做什么，错误描述如果是英文必须是小写，不需要标点结尾，采用独立的错误流进处理。

### 变量

```go
package learnvar

import "fmt"

// (name3 string, name4 string) 不写，可以是匿名变量
func test() (string, string) {
 // 初始化多个变量
 var dwj, hmm, ll = "邓文杰", "韩梅梅", "李雷"
 fmt.Printf("name1: %v\n", dwj)
 fmt.Printf("name1: %v\n", hmm)
 fmt.Printf("name1: %v\n", ll)

 //批量声明,使用一个 var 关键字，把一些变量写在一个括号()里
 var (
  hh string
  xx = "嘻嘻"
 )
 fmt.Printf("name1: %v\n", hh)
 fmt.Printf("name1: %v\n", xx)

 // 短变量声明 只能在函数内部使用，函数外部不能使用！！！！
 // 在函数内部，可以使用 := 运算符对变量进行声明和初始化
 study := "学习"
 gg := "GG"
 fmt.Printf("name1: %v\n", study)
 fmt.Printf("name1: %v\n", gg)

 // 匿名变量
 // 如果我们接收多个变量，有一些变量使用不到，可以使用下划线_表示变量名称，这种变量叫做匿名变量
 return "王小波", "李银河"
}

func LearnVar() {
 // 变量初始化语法
 // var 变量名 类型 = 表达式
 var name string = "dengwenjie"
 var age int = 23
 var gender string = "男"
 fmt.Printf("name1: %v\n", name)
 fmt.Printf("name1: %v\n", age)
 fmt.Printf("name1: %v\n", gender)

 name1, _ := test()
 fmt.Printf("name1: %v\n", name1)
}
```

### 常量

常量就是在程序编译阶段就确定下来的值，而程序在运行时则无法改变该值，在 go 程序中，常量可以是数值类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等

```go
package learnconst

import "fmt"

func LearnConst() {
	//const constantName [type] = value
	const number float64 = 31.11
	const str = "你好"
	fmt.Printf("number: %v\n", number)
	fmt.Printf("str: %v\n", str)

	//	const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		a1 = 100
		a2 = 200
		a3
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	//	iota
	// iota 比较特殊，可以被认为是一个可被编译器修改的常量，它默认开始值是 0，每调用一次加 1，
	// 遇到 const 关键字时被重置为 0 !!!!
	const a4 = iota // 0
	const a5 = iota // 0

	const (
		a6 = iota // 0
		a7 = iota // 1
		a8 = iota // 2
	)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
	fmt.Printf("a7: %v\n", a7)
	fmt.Printf("a8: %v\n", a8)

	// 使用 _ 跳过某些值
	const (
		a9 = iota // 0
		_
		a10 = iota // 2
	)
	fmt.Printf("a9: %v\n", a9)
	fmt.Printf("a10: %v\n", a10)

	//	iota 声明中间插队  和跳过的道理是一样的
	const (
		a11 = iota // 0
		a12 = 1200 // 1200
		a13 = iota // 2
	)
	fmt.Printf("a11: %v\n", a11)
	fmt.Printf("a12: %v\n", a12)
	fmt.Printf("a13: %v\n", a13)
}
```

### 数据类型

在 go 语言中，数据类型用于声明函数和变量

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要要用大数据的时候才需要申请大内存，就可以充分利用内存

* 布尔类型的值只可以是常量 true 或 false，var b bool = true
* 数字类型整型 int 和浮点型 float32、float64，go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码
* 字符串类型就是一串固定长度的字符连接起来的字符序列，go的字符串是由单个字节连接起来的，go语言的字符串的字节使用 UTF-8编码标识 Unicode 文本
* 派生类型包括：1、指针类型（Pointer），2、数组类型，3、结构体类型（struct），4、Channel类型，5、函数类型、6、切片类型，7、接口类型，8、Map类型

```go
package learndatatype

import "fmt"

func fn() {

}

func LearnDataType() {
	var name string = "邓文和"
	age := 18
	b := true
	// %T 是打印类型 \n 是换行
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)

	//	指针类型
	var a = 1
	dwj := &a
	fmt.Printf("%T\n", dwj)

	//	数组类型
	arr := [2]int{1, 2}
	fmt.Printf("%T\n", arr)

	// 切片 类型, 就是不写长度
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", s)

	// 函数类型
	fmt.Printf("%T\n", fn)
}
```

**布尔类型**

* go语言中的布尔类型有两个常量值：true 和 false，布尔类型经常用在**条件判断**语句，**循环语句**，也可以用在**逻辑表达式**中，不能使用 0和非0标识真假

```go
func main() {
	n := 1
	// 这样是不行的，不能使用 0和非0标识真假
	if n {}
}
```

**数字类型**

* go语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码
* go 也有基于架构的类型，例如：int、uint、uintptr
* 这些类型的字节都是根据程序所在的操作系统类型所决定的：
* uintptr 的长度被设定位足够存放一个指针即可
* int 和 uint 在 32 位操作系统，均使用32位（4个字节），在64位操作系统上，均使用64位（8个字节）
* go 语言没有 float 类型，只有 float32 和 float64，没有 double类型
* 与操作系统架构无关的类型都有固定的大小，并在类型的名称中就可以看出来：
* int8（-128 - 127）int16、int32、int64
* uint8（0- 255）、uint16、uint32、uint64
* float32、float64
* int 型是计算最快的一种类型
* 整型的默认值是 0，浮点型的默认值是 0.0

```go
package number

import (
	"fmt"
	"math"
	"unsafe"
)

func LearnNumber() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB% %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui8, unsafe.Sizeof(ui8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui16, unsafe.Sizeof(ui16), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui32, unsafe.Sizeof(ui32), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui64, unsafe.Sizeof(ui64), math.MinInt8, math.MaxInt8)

	var f64 float64
	var f32 float32
	fmt.Printf("%T %dB% %v~%v\n", f64, unsafe.Sizeof(f64), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", f32, unsafe.Sizeof(f32), math.MinInt8, math.MaxInt8)

	var i int
	fmt.Printf("%T %dB% %v~%v\n", i, unsafe.Sizeof(i))
}
```

**字符串类型**

* 一个 go 语言字符串是一个任意字节的常量序列。[]byte
* 字符串字面量使用双引号 ""，或者反引号``,来创建，双引号用来创建可解析的字符串，支持转义，但不能用来引用多行，反引号用来创建原生的字符串字面量，可能由多行组成，但不支持转义，并且可以包含除了反引号外其他所有字符。双引号创建可解析的字符串应用最广泛，反引号用来创建原生的字符串则多用于书写多行消息，HTML以及正则表达式。

```go
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
```

### 格式输出

```go
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
```

### 运算符

* 算数运算符
* 关系运算符
* 逻辑运算符
* 位运算符
* 赋值运算符
* ++（自增）和 --（自减）在 go 语言中是单独的语句，并不是运算符

```go
package operator

import "fmt"

func Operator() {
	a := 100
	b := 200

	c := a + b
	c = a - b
	c = a * b
	c = a / b
	fmt.Printf("%v\n", c)

	/*
	 + - * /
	 == !=
	 && || !
	 =
	 & | << >>
	*/
}
```

### 流程控制

**条件语句是用来判断给定的条件是否满足（表达式值是否为 true 或者 false），并根据判断的结果（真或假）决定执行的语句**

* if else 语句

```go
package flow

import "fmt"

func IfFlow() {
	/**
	1、不需要使用括号将条件包含起来
	2、大括号 {} 必须存在，及时只有一行语句
	3、左括号必须在 if 或 else 的同一行
	4、在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ; 进行分隔
	*/

	//初始变量可以声明在布尔表达式里面，但是作用域就在 if 语句里面了
	if age := 17; age < 18 {
		fmt.Println("不能干某些事")
	} else {
		fmt.Println("可以干某些事")
	}

	//不能使用 0 或 非 0 表示真假
	//number := 1
	//if number {
	//
	//}

	var num int
	fmt.Printf("请输入 num 值：")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Printf("\n偶数")
	} else {
		fmt.Printf("\n奇数")
	}
}
```

**switch**

```go
package flow

func fn1() {
	var num = 1
	switch num {
	case 1, 2, 3, 4, 5:
		println("工作日")
	case 6, 7:
		println("休息日")
	}
}

func fn2() {
	num1 := 2
	switch {
	case num1 > 1:
		println("大于三")
	case num1 < 1:
		println("小于一")
	}
}

func fn3() {
	num2 := 2
	switch {
	case num2 > 1:
		println("大于三")
		fallthrough
	case num2 < 1:
		println("小于一")
	}
}

func SwitchFlow() {
	/**
	1、支持多条件匹配
	2、不同的 case 之间不使用 break 分隔，默认只会执行一个 case
	3、若想要执行多个 case，需要使用 fallthrough 关键字，也可以 break 终止
	4、分支还可以使用表达式，例如：a > 10
	*/

	fn1()
	fn2()
	fn3()
}
```

**for**

* Go 语言中的 for 循环，只有 for 关键字，去除了像其他语言中的 while 和 do while
* 语法 ：for 初始语句; 条件表达式(i < 10); 结束语句 { 循环体语句 }，不用加括号
* for 循环可以通过 break、goto、return、 panic 语句强制退出循环

```go
package flow

func ForFlow() {
	for num := 0; num < 10; num++ {
		println(num)
	}

	// 类似与 while 循环
	num1 := 0
	for num1 < 5 {
		println(num1)
		num1++
	}
	println(num1)
}
```

**for range**

* Go 语言中可以使用 for range 遍历数组、切片、字符串、map 及通道（channel）。通过 for range 遍历的返回值有如下规律：
* 1、数组、切片、字符串返回索引和值
* 2、map返回键和值
* 3、通道（channel）值返回通道内的值

```go
package flow

import "fmt"

//数组
func forRangeFlow1() {
	arr := [...]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("%v:%v\n", i, v)
	}
}

//切片
func forRangeFlow2() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range s {
		fmt.Printf("%v:%v\n", i, v)
	}
}

//map
func forRangeFlow3() {
	m := make(map[string]string, 0)
	m["name"] = "邓文杰"
	m["age"] = "23"
	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
}

//字符串
func forRangeFlow4() {
	str := "hello world"
	for i, v := range str {
		fmt.Printf("%v:%c\n", i, v)
	}
}

func ForRangeFlow() {
	forRangeFlow1()
	forRangeFlow2()
	forRangeFlow3()
	forRangeFlow4()
}
```

**break**

* break语句可以结束for、switch和select的代码块
* 单独在 select 中使用 break 和不使用  break 没有区别
* 单独在表达式 switch 语句，并且没有 fallthough，使用 break 和不使用 break 没有啥区别
* 单独在表达式 switch 语句，并且有 fallthough，使用 break 能够终止 fallthough后面的case语句的执行
* 带标签的 break，可以跳出多层 select / switch 作用域。让 break 更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，没有带 break 的只能跳出当前语句块。

```go
package flow

func BreakFlow() {
	for i := 0; i < 10; i++ {
		println(i)
		if i == 3 {
			break
		}
	}

	num := 2
	switch num {
	case 1:
		println(1)
	case 2:
		println(2)
		break
	default:
		println("默认")
	}
}
```

**continue**

* continue 只能用在循环中，在 go 中只能用在 for 循环中，他可以终止本次循环，进行下一次循环，在 continue 语句后面添加标签时，表示开始标签对应的循环。

```go
package flow

func ContinueFlow() {
	for i := 0; i < 10; i++ {
		println("i:", i)
		if i == 2 {
			continue
		}
	}

label:
	for v := 0; v < 5; v++ {
		for j := 0; j < 5; j++ {
			if v == 2 && j == 2 {
				continue label
			}
			println(v, j)
		}
	}
}
```

**goto**

* goto 语句通过标签进行代码间的无条件跳转，goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。go语言中使用 goto 语句能简化一些代码的实现过程。例如双层嵌套的 for 循环要退出时

```go
package flow

func GotoFlow() {
	const num int = 1
	if num == 2 {
		println("2")
	} else {
		goto LABEL1
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				goto LABEL2
			}
		}
	}

	LABEL1:
		println("goto跳过来了")
		println("3")

LABEL2:
	println("goto2")
}
```

### 数组

* 数组是**相同数据类型**的一组数据的集合，数组一旦定义长度不能修改，数组可以通过下标（或者叫索引）来访问元素。
* 数组定义的语法如下：var nam [size]type
* name：数组名称
* size：数组长度，必须是常量
* type：数组保存元素的类型
* 省略数组长度，数组长度可以省略，使用 ... 代替，初始化值的数量自动推断
* 初始化就是给数组的元素赋值，没有初始化的数组，默认元素值都是零值，布尔类型是 false，字符串是空字符串。
* 指定索引值的方式来初始化，可以通过指定索引的方式来初始化，未指定索引的默认为类型的默认值

```go
package arr

import "fmt"

func Array() {
  // 初始化就是给数组的元素赋值，没有初始化的数组，默认元素值都是零值，布尔类型是 false，字符串是空字符串。
	var arr1 [2]int
	var arr2 [2]string
	var arr3 [2]bool
	arr1[1] = 2
	arr2[1] = "邓文杰"
	arr3[1] = true
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr4 := [3]int{1, 2, 3}
	fmt.Println(arr4)
	
  // 省略数组长度，数组长度可以省略，使用 ... 代替，初始化值的数量自动推断
	arr5 := [...]string{"李雷", "韩梅梅", "马伊琍"}
	fmt.Println(arr5)
	
  // 指定索引值的方式来初始化，可以通过指定索引的方式来初始化，未指定索引的默认为类型的默认值
	arr6 := [...]int{1: 11, 3: 133}
	fmt.Println(arr6)
}
```

**访问数组元素**

```go
package arr

import "fmt"

func Element() {
	//访问数组元素
	
	var arr1 = [3]int{1, 2, 3}
	println(arr1[0])
	println(arr1[1])
	println(arr1[2])
	//赋值
	arr1[0] = 100
	println(arr1[0])

	arr2 := [...]string{"李雷", "韩梅梅", "王小波", "李银河"}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	arr3 := [2]bool{true, false}
	for i, v := range arr3 {
		fmt.Printf("arr3[%v]: %v\n", i, v)
	}
}
```

### 切片

* 数组是固定长度，可以容纳相同数据类型的元素的集合。当长度固定时，使用还是带来一些限制，比如：我们申请的长度太大浪费内存，太小又不够用。
* 鉴于上述原因，有了切片，可以把切片理解为可变长度的数组，它底层就是用数组实现的，增加了自动扩容功能，切片（slice）是一个拥有相同类型元素的可变长度的序列。
* 声明一个切片和声明一个数组类似，只要不添加长度就行了 -> var identifier []type
* 切片是引用类型，可以使用 **make**函数来创建切片：
* var slice1 []type = make([]type, len)  -> slice := make([]type, len)
* 也可以指定容量，其中 capacity 为可选参数   make([]T, length, capacity)
* 这里 len 是数组的长度并且也是切片的初始长度
* 切片拥有自己的长度和容量，可以通过使用内置的 len() 函数求长度，使用内置的 cap() 函数求切片的容量。

```go
package slice

import "fmt"

func Slice() {
  // 切片（slice）是一个拥有相同类型元素的可变长度的序列。
	// 声明切片
	var s1 []string
	var s2 = []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)

	//使用 make 函数声明切片
	s3 := make([]string, 2)
	s3[0] = "李红"
	s3[1] = "王波"
	fmt.Printf("%v\n", s3)

	s4 := make([]string, 3, 3)
	s4[1] = "哈哈"
	fmt.Printf("%v\n", s4) // [ "哈哈" ]
	//切片长度
	fmt.Printf("%v\n", len(s4)) // 3
	//切片容量
	fmt.Printf("%v\n", cap(s4)) // 3
}
```

**切片初始化**

* make() 函数的作用 -> 用于初始化[slice](https://so.csdn.net/so/search?q=slice&spm=1001.2101.3001.7020)、chan和map

* 一个切片在末初始化之前默认为 nil，长度为 0，容量为 0

```go
package slice

import "fmt"

func InitSlice() {
	//	初始化切片
	s5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", s5)

	//使用数组初始化
	arr := [...]int{1, 2, 3, 4}
	s6 := arr[:]
	fmt.Printf("%v\n", s6)

	//截取
	s7 := []int{1, 2, 3, 4, 5, 6}
	s8 := s7[:3]
	s9 := s7[3:]
	s10 := s7[2:5]
	s11 := s7[:]
	fmt.Printf("%v\n", s8)
	fmt.Printf("%v\n", s9)
	fmt.Printf("%v\n", s10)
	fmt.Printf("%v\n", s11)
}
```

**切片增删改查**

* 切片是一个动态数组，可是使用 append() 函数添加元素，go 语言中并没有删除切片元素的专用方法，可以使用切片本身的特性来删除元素，由于切片是引用类型，通过赋值的方式，会修改原有内容，go 提供了 copy() 函数来拷贝切片

```go
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
```

### Map

* map 是一种 key: value 键值对的数据结构容器，map 内部实现是哈希表（hash）。map 最重要的一点是通过 key 来快速检索数据，key 类似与索引，指向数据的值，map 是引用类型
* map 的语法格式 var m map[key_data_type]value_data_type
* 使用 make 函数初始化，m := make(map[key_data_type]value_data_type)
* 判断某个键是否存在，v, ok := m[key]，如果 ok 为 true，存在，反之

```go
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
}
```

### 函数

* 函数在 go 语言中手**一级公民**，我们把所有的功能单元都定义在函数中，可以重复使用，函数包含函数的名称，参数列表和返回值类型，这些构成了函数的签名
* go 语言中有 3 种函数：普通函数、匿名函数、方法（定义在 struct 上的函数）
* go 语言中不允许函数重载（overload），也就是说不允许函数同名
* go 语言中的函数不能嵌套函数，但可以嵌套匿名函数
* 函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数
* 函数可以作为参数传递给另一个函数
* 函数的返回值可以是一个函数
* 函数调用的时候，如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数
* 函数参数可以没有名称
* 函数在使用之前必须定义，可以调用函数来完成某个任务，函数可以重复调用，从而达到代码重用

```go
package function

import "fmt"

func f1(a int, b int) (c int) {
	c = a + b
	return c
}

func f2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LFunction() {
	var c = f1(1, 2)
	fmt.Printf("%v\n", c)

	var d = f2(1, 2)
	fmt.Printf("%v\n", d)
}
```

* Go 语言中经常会使用其中一个返回值作为函数是否执行成功，是否有错误信息的判断条件，例如：return value, exists、return value, ok、return value, err 等
* 当函数的返回值过多时，应该将这些返回值收集到容器中，然后以返回容器的方式去返回。例如，同类型的返回值可以放进 slice中，不同类型的返回值可以放进 map 中

```go
package function

import "fmt"

func f3() {
	fmt.Printf("%v\n", "f1 没有返回值")
}

func f4() string {
	return "有一个返回值"
}

func f5() (string, int) {
	return "多个返回值", 2
}

func f6() (name string, age int) {
	name = "邓文杰"
	age = 23
	return name, age
}

func f7() (name string, age int) {
	name = "余华"
	age = 18
	return // 等价于 return name, age
}

func LFuncReturn() {
	f3()
	str := f4()
	fmt.Printf("%v\n", str)
	str1, n := f5()
	fmt.Printf("%v\n%v\n", str1, n)
	name, age := f6()
	fmt.Printf("%v\n%v\n", name, age)
	name1, age1 := f7()
	fmt.Printf("%v\n%v\n", name1, age1)
}
```

**go 函数的参数**

* 函数可以有 0 或多个参数，参数需要指定数据类型
* 声明函数时的参数列表叫做形参，调用时传递的参数叫做实参
* go 语言是通过**传值的方式传参**，意味着传递给函数的是拷贝后的副本，所以函数内部访问，修改的也是这个副本
* go 语言可以使用可变参数，有时候并不能确定参数的个数，可以使用变长参数，可以在函数定义语句的参数部分使用 args...type的方式。这时会将...代表的参数全部保存到一个名为 args 的 slice 中，注意这些参数的数据类型都是 type
* **map、slice、interface、channel这些数据类型本身就是指针类型的，所以就算是拷贝值也是拷贝的指针，拷贝后的参数仍然指向底层数据结构，所以修改它们可能会影响外部数据结构的值**

```go
package function

import "fmt"

func f8(a int) {
	a = 200
}

func f9(args ...int) {
	fmt.Printf("%v\n%T", args, args)
}

func f10(name string, age int, args ...string) {
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", age)

	for i, v := range args {
		fmt.Printf("%v\n%v\n", i, v)
	}
}

func f11(s1 []int) {
	s1[0] = 20000
}

func LFuncParams() {
	i1 := 100
	f8(i1)
	fmt.Printf("%v\n", i1)

	f9(1, 2, 3, 4, 5)

	f10("邓文杰", 23, "韩梅梅", "李雷")

	var s = []int{1, 2, 3, 4, 5}
	f11(s)
	fmt.Printf("%v\n", s)
}
```

**函数类型与函数变量**

* 可以使用 type 关键字来定义一个函数类型，

* ```go
  type fun func(int, int) int
  ```

* 上面语句定义了一个 fun 函数类型，它是一个函数类型，这种函数接收两个 int 类型的参数并且返回一个 int 类型的返回值

```go
package function

import "fmt"

type fun func(int, int) int

func f14(a int, b int) int {
	return a + b
}

func f16(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LFuncType() {
	var f15 fun
	f15 = f14
	num := f15(1, 2)
	fmt.Printf("%v\n", num)

	var f17 fun
	f17 = f16
	max := f17(1, 2)
	fmt.Printf("%v\n", max)
}
```

**高阶函数**

* go语言中，函数可以作为参数，传递给另一个函数，函数可以作为返回值，另外一个函数返回一个函数

```go
package function

import "fmt"

func ff1(name string, callback func(str string)) {
	callback(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

type fn func(int2 int, int3 int) int
func cal(o string) fn {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func LfuncG() {
	ff1("邓文杰", func(str string) {
		fmt.Printf("%v\n", str) // 邓文杰
	})

	n := cal("+")(2, 3)
	fmt.Printf("%v\n", n)
	
	n1 := cal("-")(2, 1)
	fmt.Printf("%v\n", n1)
}
```

**匿名函数**

* go 语言中函数不能嵌套，但是在函数内部可以定义匿名函数，匿名函数就是，没有名称的函数

* ```
  func (参数列表)(返回值)
  ```

```go
package function

import "fmt"

func LFuncN() {
	//匿名函数
	f := func() {
		fmt.Printf("%v\n", "我是匿名函数")
	}
	f()

	//立即执行函数、自己调用自己
	(func() {
		fmt.Printf("%v\n", "我是立即执行函数")
	})()
}
```

**闭包**

* 闭包可以理解成定义在一个函数内部的函数，在本质上，闭包是将函数内部和函数外部连接起来的桥梁，或者说是函数和其引用环境的组合体
* 闭包指的是一个函数和其相关的引用环境组合而成的实体，简单来说，**闭包 = 函数 + 引用环境**（重要）。这个变量不会销毁

```go
package function

import "fmt"

// 变量 numFn 是一个函数并且它引用了其外部作用域中的 x 变量，此时 numFn 就是第一个闭包，在 numFn  的生命周期内，变量 x 也一直有效
func f20() func(int) int {
	var x int
	//用到了闭包 -> 闭包 = 函数 + 引用环境
	return func(num int) int {
		x += num
		return x
	}
}

type fn1 func(int) int

func cal1(base int) (fn1, fn1) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(b int) int {
		base -= b
		return base
	}
	return add, sub
}

func LFuncClosure() {
	//numFn := f20()
	//n := numFn(10)
	//fmt.Printf("%v\n", n) // 10
	//n = numFn(20)
	//fmt.Printf("%v\n", n) // 30
	//n = numFn(30)
	//fmt.Printf("%v\n", n) // 60
	//
	//numFn1 := f20()
	//n2 := numFn1(100)
	//fmt.Printf("%v\n", n2) // 100
	//n2 = numFn1(200)
	//fmt.Printf("%v\n", n2) // 300

	add, sub := cal1(100)
	res := add(1)
	fmt.Printf("%v\n", res) // 101
	res = sub(2)
	fmt.Printf("%v\n", res) // 99
	res = add(3)
	fmt.Printf("%v\n", res) // 102
	res = sub(4)
	fmt.Printf("%v\n", res) // 98
}
```

**递归**

* 函数内部调用函数自身的函数称为递归函数
* 使用递归函数最重要的三点：
* 1、递归就是自己调用自己
* 2、必须先定义函数的退出条件，没有退出条件，递归将成为死循环
* 3、go 语言递归函数很可能会产生一大堆的 goroutine，也很可能会痴线栈空间内存溢出问题

```go
package function

import "fmt"

func f21(n int) int {
	if n == 1 {
		return 1
	}
	// 函数调用栈，后进先出，上面的函数执行完先销毁
	return n * f21(n-1)
}

func LFuncRuc() {
	i := f21(5)
	fmt.Printf("%v\n", i) // 120
}
```

**defer语句**

* go 语言中的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 定义的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行
* defer 特性：
* 1、关键字 defer 用于注册延迟调用
* 2、这些调用直到 return 前才被执行，因此可以用来做资源清理
* 多个 defer 语句，按先进后出的方式执行 stack
* defer 语句中的变量，在 defer 声明时就决定了
* defer 用途：
* 1、关闭文件句柄
* 2、锁资源释放
* 3、数据库连接释放

```go
package function

func LFuncDefer() {
	print("start...")
	defer print("我是 defer1")
	defer print("我是 defer2")
	defer print("我是 defer3")
	print("end...")
}
```

**init函数**

* go 语言有一个特殊的函数 init 函数，先于 main 函数执行，实现包级别的一些初始化操作
* init 函数的主要特点
* 1、init 函数先于 main 函数自动执行，不能被其他函数调用
* 2、init 函数没有输入参数、返回值
* 3、每个包可以有多个 init 函数
* 4、包的每个源文件也可以有多个 init 函数，这点比较特殊
* 5、同一个包的 init 执行顺序，golang 没有明确定义，编程时要注意程序不要依赖这个执行顺序
* 6、不同包的 init 函数按照包导入的依赖关系决定执行顺序
* golang 初始化顺序：**变量初始化 -> init() -> main()**

```go
package main

import "go_ww/function"

var dwj = initVar()

func init() {
	println("我是初始化函数 init 系统调用的") // 第二
}

func initVar() string {
	println("变量最先") // 第一
	return "邓文杰"
}

func main() {
	function.LFuncDefer() // 最后
}
```

### 指针

* go 语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无需拷贝数据。

* 类型指针不能进行偏移和运算

* go 语言中的指针操作非常简单，只需要记住两个符合：&取地址和*根据地址取值

* 指针地址和指针类型：每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置，go 语言中使用 & 字符放在变量前面对变量进行**取地址**操作，go 语言中的**值类型（int、float、bool、string、array、struct）**都有对应的指针类型，如：* int、* int64、*string等。

* 指针语法：一个指针变量指向了一个值的内存地址。也就是我们声明了一个指针之后，可以像变量赋值一样，把一个值的内存地址放入到指针当中。

* ```go
  var var_name *var_type
  ```

* var_type：指针类型

* var_name：指针变量名

* *：用于指定变量是作为一个指针

```go
package pointer

import "fmt"

func LPointer() {
	//指针变量
	var p *int
	fmt.Printf("%v\n", p) // nil
	num := 100
	//拿到 num 的地址赋值给 p
	p = &num
	fmt.Printf("%v\n", p)  // 0x1400012c010
	fmt.Printf("%T\n", p)  // *int
	fmt.Printf("%v\n", *p) // 100
}
```

**数组的指针**

* 定义数组

* ```go
  var ptr [2]*int // 表示数组里面的元素的类型是指针类型
  ```

```go
package pointer

import "fmt"

func LPointerArr() {
	//表示数组里面的元素的类型是指针类型
	var p1 [3]*int
	arr := [3]int{1, 2, 3}

	fmt.Printf("%v\n", p1) // [<nil> <nil> <nil>]
	for i := 0; i < len(arr); i++ {
		p1[i] = &arr[i]
	}
	fmt.Printf("%v\n", p1) // [0x140000ac000 0x140000ac008 0x140000ac010]

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\n", *p1[i])
	}
}
```

### 类型定义和类型别名

* 在介绍**结构体**之前，先来看看什么是类型定义和类型别名

* 类型定义的语法

* ```go
  type NewType Type
  ```

* 类型别名

* ```go
  type ww = string
  ```

* 类型定义和类型别名的区别：

* 1、类型定义相当于定义了一个全新的类型，与之前的类型不同，但是类型别名并没有定义一个新的类型，而是使用一个别名来替代之前的类型

* 2、类型别名只会在代码中存在，在编译完成之后并不会存在该别名

* 因为类型别名和原来的类型是一致的，所以原来类型所拥有的方法，类型别名中也可以调用，但是如果是重新定义的一个类型，那么不可以调用之前的任何方法。

```go
package _type

import "fmt"

func LTypeDefine() {
	//类型定义
	type NewInt int
	var num NewInt = 100

	fmt.Printf("%v%T\n", num, num) // 100 _type.NewInt

	//	类型别名
	type ww = string
	var str ww
	str = "ww"
	fmt.Printf("%v%T\n", str)
}
```

### 结构体

* go 语言中没有面向对象的概念，但是可以使用结构体来实现， 面向对象编程的一些特性：继承、组合等特性

* go 语言结构体的定义，结构体的定义和类型定义差不多，只不过多了一个 struct 关键字

* ```go
  type struct_variable_type struct {
  	name string
  	age int
  }
  ```

* 同类型的可以放在一起

* ```go
  type Person struct {
  	name, email string
  	age int
  }
  ```

```go
package _struct

import "fmt"

type Person struct {
	name, email string
	age         int
}

func LStruct() {
	var s Person
	
	fmt.Printf("%v\n", s) // {  0}

	s.name = "邓文杰"
	s.age = 23
	s.email = "coderdwj@163.com"

	fmt.Printf("%v\n", s) // {邓文杰 coderdwj@163.com 23}

	//	匿名结构体
	var dwj struct {
		age   int
		hobby []string
	}
	dwj.age = 18
	dwj.hobby = []string{"社会人", "睡觉"}
	fmt.Printf("%v\n", dwj) // {18 [社会人 睡觉]}
}
```

**结构体初始化**

* 为初始化的结构体，成员都是零值 int 0, float 0.0，bool false，string ""，nil nil

```go
package _struct

import "fmt"

type Person1 struct {
	id, age     int
	name, email string
	hobby       []string
}

func LStructInit() {
	//结构体初始化
	var p = Person1{
		id:    2023,
		age:   23,
		name:  "邓文杰",
		email: "coderdwj@163.com",
		hobby: []string{"写博客", "听赵雷的歌曲"},
	}

	fmt.Printf("%v\n", p) // {2023 23 邓文杰 coderdwj@163.com [写博客 听赵雷的歌曲]}

	//使用列表的方式进行初始化
	p1 := Person1{
		2023,
		23,
		"邓文杰",
		"coderdwj@163.com",
		[]string{"民谣"},
	}
	fmt.Printf("%v\n", p1) // {2023 23 邓文杰 coderdwj@163.com [民谣]}

	//	部分初始化
	p2 := Person1{
		name: "邓文杰",
		age:  23,
	}
	fmt.Printf("%v\n", p2) // {0 23 邓文杰  []}
}
```

* 列表的方式初始化必须初始化所有字段，必须和字段在结构体中的声明顺序一致，该方式不能混用

**结构体指针**

* 结构体指针和普通的变量指针相同
* 可以使用 new 关键字创建结构体指针

```go
package _struct

import "fmt"

func LStructPointer() {
	p := Person1{
		name: "Deng",
		age:  23,
	}
	//结构体指针
	var pP *Person1
	pP = &p
	fmt.Printf("%v\n", pP)  // &{0 23 Deng  []}
	fmt.Printf("%p\n", pP)  // 0x14000100050
	fmt.Printf("%v\n", *pP) // {0 23 Deng  []}

	//使用 new 关键字创建结构体指针
	var pP1 = new(Person1)
	pP1 = &p
	//(*pP1).id = 101
	pP1.id = 101
	pP1.hobby = []string{"吃饭"}
	fmt.Printf("%p\n", pP1)  // 0x1400008c050
	fmt.Printf("%v\n", *pP1) // {101 23 Deng  [吃饭]}
}
```

**结构体作为函数的参数**

* go 结构体可以像普通变量一样，作为函数的参数，传递给函数，这里分两种情况（结构体是值类型）
* 1、直接传递结构体，这是一个副本，在函数内部不会改变外面结构体内容
* 2、传递结构体指针，这时在函数内部，能够改变外部结构体内容

```go
package _struct

import "fmt"

type Person2 struct {
	name string
	age  int
}

//1、直接传递结构体
func test1(p Person2) {
	p.age = 101
	p.name = "韩梅梅"
	fmt.Printf("%v\n", p) // {韩梅梅 101}
}

//2、传递的是结构体指针 会修改原来的
func test2(pP *Person2) {
	pP.name = "李银河"
	pP.age = 11

	fmt.Printf("%v\n", *pP) // {李银河 11}
}

func LStructParams() {
	//结构体作为参数传递给函数
	p := Person2{
		name: "李雷",
		age:  18,
	}
	test1(p)
	fmt.Printf("%v\n", p) // {李雷 18}，不会改变，因为是值类型

	p1 := Person2{
		name: "王小波",
		age:  10,
	}
	pP := &p1
	test2(pP)              // test2(&p1)
	fmt.Printf("%v\n", p1) // {李银河 11}
}
```

**结构体嵌套**

* go语言没有面向对象编程的思想，也没有继承关系，但是可以通过结构体嵌套来实现这种效果

```go
package _struct

import "fmt"

func LStructQ() {
	//结构体嵌套
	type Dog struct {
		name  string
		color string
	}

	type Person struct {
		name string
		age  int
		dog  Dog
	}

	p := Person{
		name: "莫言",
		age:  18,
		dog: Dog{
			"小花",
			"红色",
		},
	}
	fmt.Printf("%v\n", p) // {莫言 18 {小花 红色}}
}
```

### go 语言方法

* go 语言没有面向对象的特性，也没有类对象的概念，但是，可以使用结构体来模拟这些特性，我们都知道面向对象里面有类方法等概念，我们也可以声明一些方法，属于某个结构体

* go 中的方法，是一种特殊的函数，定义于 struct 之上（与 struct 关联、绑定），被称为 struct 的接受者（receiver）。通俗的讲，方法就是有接收者的函数。

* ```go
  type Person struct {
  	name string
  }
  func (p Person) eat() {}
  func (p Person) say() {}
  ```

* p：接收该方法的结构体

```go
package _struct

import "fmt"

type Person3 struct {
	name string
	say  func()
}

// 结构体方法，struct 的接收者
func (p Person3) eat() {
	fmt.Printf("%v\n", p.name)
	fmt.Printf("%v\n", "我是Person3 结构体的方法")
}

func (p Person3) login(username string, password string) bool {
	if username == "邓文杰" && password == "123456" {
		fmt.Printf("%v\n", p.name)
		return true
	}
	return false
}

func LStructMethod() {
	p1 := Person3{
		name: "邓文杰",
		say: func() {
			// 这样写拿不到属性，?目前不知道
			println("进")
		},
	}
	p1.eat()
	p1.say()

	isT := p1.login("邓文杰", "123456")
	if isT {
		println("登录成功")
	} else {
		println("用户名或密码错误")
	}
}
```

* Go 语言方法的注意事项
* 1、方法的 receiver type 并非一定要是 struct 类型，type 定义的类型别名、slice、map、channel、func 类型等都可以
* 2、struct 结合它的方法就等价于面向对象中的类。只不过 struct 可以和它的方法分开，并非一定要属于同一个文件，但必须属于同一个包。
* 3、方法有两种接收类型：（T Type）和（T *Type）,它们之间有区别
* 4、方法就是函数，所以go中没有方法重载（overload）的说法，也就是说同一个类型中的所有方法名必须都唯一。
* 5、如果 receiver 是一个指针类型，则会自动解除引用->(*p).name
* 6、方法和 type 是分开的，意味着实例的行为和数据存储是分开的，但是它们通过 receiver 建立起关联关系。

**go方法接收者类型**

* 结构体实例，有值类型和指针类型，那么方法的接收者是结构体，那么也有值类型和指针类型。区别就是接收者是否复制结构体副本。值类型复制，指针类型不复制。

```go
package _struct

import "fmt"

type Person4 struct {
	name string
}

// 值类型
func (p Person4) who() {
	p.name = "不会修改"
}

// 指针类型
func (p *Person4) who2() {
	// p.name 自动解引用
	// (*p).name = "会修改"
	p.name = "会修改"
}

func test(p1 Person4) {
	p1.name = "不会修改"
}

func test3(p1 *Person4) {
	p1.name = "会修改哦"
}

func LStructMethodArgs() {
	p := Person4{
		"邓文杰",
	}
	p.who()
	fmt.Printf("%v\n", p) // {邓文杰}
	p.who2()
	fmt.Printf("%v\n", p) // {会修改}

	p1 := Person4{
		"莫言",
	}
	test(p1)
	fmt.Printf("%v\n", p1) // {莫言}

	test3(&p1)
	fmt.Printf("%v\n", p1) // {会修改哦}
}
```

### interface

* go 语言中的接口，是一种新的类型定义，它把所有的具有共性的方法定义在一起，**任何其他类型只要实现了这些方法就是实现了这个接口**。

```go
package _interface

// USB 定义接口
type USB interface {
	read()
	write()
}

type Computer struct {
	Name string
}

type M struct {
	Name string
}

// 方法
func (m M) read() {
	println("我是m read")
}

func (m M) write() {
	println("我是m read")
}

func (c Computer) read() {
	println("我是c read")
}

func (c Computer) write() {
	println("我是c write")
}

func LInterface() {
	c := Computer{
		Name: "电脑",
	}
	var u USB = c
	u.read()
	u.write()

	m := M{
		Name: "手机",
	}

	test(c)
	test(m)
}

func test(u USB) {
	u.read()
	u.write()
}
```

**go 接口值类型接收者和指针类型接收者**

* 本质上和方法的值类型接收者和指针类型接收者的思考方法是一样的

```go
package _interface

import "fmt"

type Pet interface {
	eat()
	say()
}

type Cat struct {
	Name string
}

//func (c Cat) eat() {
//	c.Name = "tom"
//}
//
//func (c Cat) say() {
//	c.Name = "kk"
//}

func (c Cat) eat() {
	c.Name = "tom"
}

func (c *Cat) say() {
	c.Name = "kk"
}

func LInterfaceReceiverType() {
	c := Cat{
		Name: "mm",
	}
	//var p Pet = c
	//p.eat()
	//p.say()
	//
	//fmt.Printf("%v\n", c) // {mm}

	var p Pet = &c
	p.eat()
	p.say()

	fmt.Printf("%v\n", c) // {kk}
}
```

* 当一个变量为接口类型，1.实现类的方法接收者都是值类型，可以把具体实现类的值或者指针赋值给这个变量，然后调用实现方法。2. 当实现类中只要有一个方法接收者定义为指针，只能把指针赋值给这个变量，调用方法。

**接口和类型的关系**

* 一个类型可以实现多个接口
* 多个类型可以实现同一个接口（多态）
* 结构体实现接口里面的方法

```go
package _interface

import "fmt"

// Pet1 一个类型可以实现多个接口
type Pet1 interface {
	p1()
}
type Pet2 interface {
	p2()
}

// I1 一个接口可以被多个类型实现
type I1 interface {
	say()
}

type Ppp struct {
}

type S1 struct {
}
type S2 struct {
}

func (p Ppp) p1() {
	fmt.Printf("%v\n", "11")
}

func (p Ppp) p2() {
	fmt.Printf("%v\n", "22")
}

func (s S1) say() {
	println("斤1")
}
func (s S2) say() {
	println("进2")
}

func LInterfaceType() {
	//var pp Pet1 = Ppp{}
	//var ppp Pet2 = Ppp{}
	//pp.p1()
	//ppp.p2()

	var i I1
	i = S1{}
	i.say()
	i = S2{}
	i.say()
}
```

**接口嵌套**

```go
package _interface

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
	Name string
}

func (f Fish) fly() {
	println("我是飞")
}

func (f Fish) swim() {
	println("我是鱼")
	f.Name = "我啥谁"
}

func LInterfaceCompose() {
	// 接口类型
	var f FlyFish
	f = &Fish{}
	f.swim()
	f.fly()
}
```

**go 通过接口实现 ocp 设计原则**

* 面向对象的可复用设计的第一块基石，便是所谓的“开-闭”原则（Open-Closed Principle，常缩写为 ocp）。虽然，go 不是面向对象语言，但是也可以模拟实现这个原则，对 **扩展**是开放的，对**修改**是关闭的

```go
package _interface

import "fmt"

type Animal interface {
	eat(string) string
	say()
}

type Dog struct {
	Name string
}

type Cat1 struct {
	Name string
}

type Person struct {
	Name string
}

func (d Dog) eat(e string) string {
	d.Name = "狗狗"
	println("我是狗狗 eat")
	return e
}
func (d Dog) say() {
	println("我是狗狗 say")
}

func (c Cat1) eat(string2 string) string {
	println("我是猫猫 eat")
	return string2
}

func (c Cat1) say() {
	println("我是猫猫 say")
}

func (p Person) pp(a Animal) {
	fmt.Printf("%v\n%T\n", a, a)
	a.eat("狗狗")
	a.say()
}

func LInterfaceOCP() {
	p := Person{}
	p.pp(new(Dog))
	p.pp(new(Cat1))
}
```

**go 模拟 oop 的属性和方法**

* go 没有面向对象的概念，也没有封装的概念，但是可以通过结构体 struct 和函数绑定来实现 oop 的属性和方法等特性，接收者 receiver 方法。

```go
package _interface

type Person2 struct {
	// 大写外面包可以访问，小写不可以
	name string
	age  int
}

func (p Person2) eat() {
	println("我可以吃")
}
func (p Person2) work() {
	println("我可以工作")
}

func LStructOOP() {
	p := Person2{
		name: "dengWJ",
		age:  18,
	}
	p.eat()
	p.work()
}
```

### 模拟继承

* golang 本质上没有 oop 的概念，也没有继承的概念，但是可以通过结构体嵌套实现这个特性

```go
package _interface

import "fmt"

type An struct {
	name string
	age  int
}

func (a An) eat(string2 string) {
	println(string2 + "实现继承")
}
func (a An) sleep(string2 string) {
	println(string2 + "实现继承 sleep")
}

type Cat3 struct {
	color string
	An    // 实现了继承
}

type Dog3 struct {
	ddd string
	An
}

// LStructExtends 实现继承
func LStructExtends() {
	c := Cat3{
		"白色",
		An{name: "小猫", age: 3},
	}
	fmt.Printf("%v\n", c.color)
	c.eat("猫猫")
	c.sleep("猫猫")

	d := Dog3{
		"ddd",
		An{
			name: "小狗",
			age:  4,
		},
	}
	fmt.Printf("%v\n", d.ddd)
	d.eat("狗狗")
	d.sleep("狗狗")
}
```

### go 模拟构造函数

* 使用函数模拟构造函数

```go
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
```

### golang 包

* 包可以区分命名空间（一个文件夹中不能有两个同名文件），也可以更好的管理项目。go 中创建一个包，一般是创建一个文件夹，在该文件夹里面的 go 文件中，使用 package 关键字声明包名称，通常，文件夹名称和包名称相同，并且，同一个文件下面只有一个包。

* 创建包：

* 1、创建一个名为 dao 的文件夹

* 2、创建一个 dao.go 文件

* 3、在该文件中声明包

* ```go
  package dao
  import "fmt"
  func Test1() {}
  ```

* 导入包，要使用某个包下面的变量或者方法，需要导入该包，导入包时，要导入从 `GOPATH`开始的包路径，例如，在 service.go 中导入 dao 包。

* 包注意事项

* 1、一个文件夹下只能有一个 package

  * `import`后面的其实是 `GOPATH`开始的相对目录路径，包括最后一段，但由于一个目录下只能有一个 package，所以 import 一个路径就等于是 import 了这个路径下的包。
  * 注意，这里指的是“直接包含”的 go 文件，如果有子目录，那么子目录的父目录是完全两个包。

* 比如实现了一个计数器 package，名叫 calc，位于 calc 目录下，但又想给别人一个使用范例，于是在 calc下可以建个 example 子目录（calc/example/），这个子目录里有个 example.go（calc/example/example.go）。此时，example.go可以是 main 包，里面还可以有个 main 函数。

* 一个 package 的文件不能在多个文件夹下

  * 如果多个文件夹下有重名的 package，它们其实是彼此无关的 package
  * 如果一个 go 文件需要同时使用不同目录下的同名 package，需要在 improt 这些目录时为每个目录指定一个 package 的别名。

go mod tidy 下载依赖

### golang 包管理工具 go module

* go module 是 golang 1.11 新加的特性，用来管理模块中包的依赖关系
* go mod 使用方法
* 初始化模块：go mod init <项目模块名称>
* 依赖关系处理，根据 go.mod 文件：go mod tidy
* 将依赖包复制到项目下的 vendor 目录：go mod vendor，如果包被屏蔽（墙），可以使用这个命令，随后使用 go build -mod=vendor 编译
* 显示依赖关系：go list -m all
* 显示详细依赖关系：go list -m -json all
* 下载依赖 go mod download [path@version]，[path@version]是非必须的

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"dengwj.vip/module/services"
)

func main() {
	println("hello world")

	services.TestUser()
	services.TestLogin()
	fmt.Printf("%v\n", "你好")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

