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

