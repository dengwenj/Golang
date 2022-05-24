# Golang 知识点笔记
## Golang 设计理念：一种事情有且只有一种方式完成

### 变量

- 变量表示内存中的一个存储区域
- 该区域有自己的名称（变量名）和类型（数据类型）
- 变量如果没有赋初始值，会有默认值，int 默认值为 0 string 默认值为空串。小数默认值为 0
- 同一个作用域下不能重名

```go
package main

import "fmt"

// 全局的变量
var uName = "dengwenjie"
var age = 22
var sex = "男"

func main() {
	/*
		变量表示内存中的一个存储区域
		该区域有自己的名称（变量名）和类型（数据类型）
		变量如果没有赋初始值，会有默认值，int 默认值为 0 string 默认值为空串。小数默认值为 0
		同一个作用域下不能重名
	*/
	var hh string
	var num int = 1
	// 一次声明多个变量的方式
	var num1, num2, num3 = 11, 22, 33
	// := 方式声明
	name := "dengwj"
	hh = "ww"
	// 重新赋值
	name = "dwj"

	//也可以
	var (
		name1 = uName
		age1  = age
		sex1  = sex
	)

	var str1 = "dwj"
	var str2 = "ww"
	var n1 = 1
	var n2 = 2
	fmt.Println(str1 + str2)
	fmt.Println(n1 + n2)

	fmt.Println(num, num1, num2, num3, name, hh)
	fmt.Println(name1, age1, sex1)
}
```



### 数据类型

- **数值类型**

```go
package main

import "fmt"

func main() {
  /*
      int 有符合
      uint 无符合
      rune 有符合 相当于 int32
      byte 无符合 相当于 uint8
  */
  fmt.Println()

  // 整形的使用细节
  var n1 = 100
  // 查看 n1 的数据类型
  // fmt.Printf() 可以用于做格式化输出
  fmt.Printf("n1 的数据类型是 %T\n", n1)

  // 如何在程序中查看某个变量的占用字节大小和数据类型
  var n2 int8 = 22
  // unsafe.Sizeof(n2) 是 unsafe 包的一个函数，可以返回 n2 变量占用的字节数
  fmt.Printf("n2 的数据类型是 %T n2 占用的字节数是 %d\n", n2, unsafe.Sizeof(n2))

  // go 程序中整型变量在使用时，遵守保小不保大原则
  // 0 ~ 255
  var age byte = 17
  fmt.Println(age)
}
```

- **浮点类型**

```go
package main

import "fmt"

func main() { 
  num := 124.11111
  fmt.Printf("num 的数据类型数 %T", num) // float64 默认
  
  // 有固定的的范围和字段长度，不受具体 os 的影响
  // 科学计数法形式：5.1234e2 = 512.34 即 5.12 * 10的二次方
  // 通常情况下，应该使用 float64
```

- **字符类型**

  **字符串就是一串固定长度的字符连接起来的字符序列。go的字符串是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成的，而 go 的字符串不同，他是由字节组成的**

  Go 中没有专门的字符类型，如果要存储单个字符（字母），一般使用 byte 来保存

  - 字符常量是用单引号括起来的单个字符，var b byte = 'a' var d int = '邓'
  - go 中允许使用转义字符 \ 来将其后的字符传变为特殊字符型常量
  - go 语言的字符使用 UTF-8 编码 英文字母 1 个字节，汉字 3 个字节
  - **go 中，字符的本质就是一个整数**，直接输出时，是该字符对应的 UTF-8 编码的码值
  - 可以直接给某个变量赋一个数字，然后按格式化输出时 %c,会输出该数字对应的 Unicode 字符
  - 字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码

  字符型存储到计算机中，需要将字符对应的码值（整数）找出来

  存储：字符->对应码值->二进制->存储

  读取：二进制->码值->字符->读取

  字符和码值的对应关系是通过字符编码表决定的(规定好的)

  go 语言的编码都统一成了 utf-8，很统一，没有编码乱码的困扰

```go
package main

import "fmt"

func main() {
  // 当要存储字符是选用 byte
  // Go 中没有专门的字符类型，如果要存储单个字符（字母），一般使用 byte 来保
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
```

- **布尔类型**

```go
package main

import (
  "fmt"
  "unsafe"
)

func main() {
  b := true
  // bool 类型占用存储空间是 1 个字节
  // bool 类型只能去 true 和 false
  
  fmt.Println("b 的占用空间是", unsafe.Sizeof(b)) // 1
}
```

- **字符串类型**

```go
package main

import "fmt"

func main() {
 // 字符串就是一串固定长度的字符连接起来的字符序列
 // go 的字符串是由单个字节连接起来的
 // go 语言的字符串的字节使用 utf-8 编码标识 Unicode 文本
 str := "你好世界 hello world"
 fmt.Println(str)
  
 str1 := `
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
	`
	fmt.Println(str1)

	str2 := "hello12332112345677887921" + "hello12332112345677887921" +
		"hello12332112345677887921" + "hello12332112345677887921" +
		"hello12332112345677887921"
	fmt.Println(str2)
}
```

- **基本数据类型默认值**

  在go中所有的数据类型都有一个默认值

```go
    // 基本数据类型的默认值
	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false
	fmt.Println(a, b, c, d) // 0 0 "" false
```

- **基本数据类型的转换**
  - go 在不同类型的变量之间赋值时需要显示转换，也就是说 go 中数据类型不能自动转换
  - 基本语法 表达式 T(v) 将值 v 转换为类型 T T: 就是数据类型，比如 int32 int64 v: 就是需要转换的变量
  - go 中，数据类型的转换可以是从表示范围小 -> 表示范围大，也可以范围大 -> 范围小
  - 被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
  - 在转化中，比如将 int64 转成 int8 【-128 ~ 127】编译时不会报错，只是转换的结果是按溢出处理，和希望的结果不一样

```go
package main

import "fmt"

func main() {
	/*
		go 在不同类型的变量之间赋值时需要显示转换，也就是说 go 中数据类型不能自动转换
		基本语法
			表达式 T(v) 将值 v 转换为类型 T
			T: 就是数据类型，比如 int32 int64
			v: 就是需要转换的变量
	*/
	a := 100
	var b float64 = float64(a)
	fmt.Println(b)
	var c int8 = int8(a)
	fmt.Println(c)

	/*
		细节说明
			1 go 中，数据类型的转换可以是从表示范围小 -> 表示范围大，也可以范围大 -> 范围小
			2 被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
			3 在转化中，比如将 int64 转成 int8 【-128 ~ 127】编译时不会报错，只是转换的结果是按溢出处理，和希望的结果不一样
	*/
	var num int64 = 134
	var num1 int8
	num1 = int8(num)
	fmt.Println(num1) // 不是 134 因为溢出了

	var n1 int32 = 12
	var n2 int64
	var n3 int8
	// n2 = n1 + 10 // 类型不一样
	// n3 = n1 + 10 // 类型不一样
	n2 = int64(n1) + 10
	n3 = int8(n1) + 10
	fmt.Println(n2, n3)

	// var n4 int32 = 22
	// var n5 int8
	// var n6 int8
	// n5 = int8(n4) + 127 // 溢出了 127 + 22 大于了 127 结果不是 149
	// n6 = int8(n4) + 128 // 128 直接溢出了
}
```

* **基本数据类型转 string**
  * 方式1：fmt.Sprintf("%参数", 表达式)
  * fmt.Sprintf() 会返回转换后的字符串
  * 方式2： strconv 方式

```go
package main

import (
 "fmt"
 "strconv"
)

func main() {
 /*
  基本类型转 string 类型
  方式1：fmt.Sprintf("%参数", 表达式)
  fmt.Sprintf() 会返回转换后的字符串
 */
 var num1 int32 = 100
 var num2 float64 = 12.09
 var b bool = true
 var str string
 // fmt.Sprintf 方式
 str = fmt.Sprintf("%d", num1)
 fmt.Printf("%T %q\n", str, str)

 str = fmt.Sprintf("%f", num2)
 fmt.Printf("%T %q\n", str, str)

 str = fmt.Sprintf("%t", b)
 fmt.Printf("%T %q\n", str, str)

 // strconv 方式
 var num3 int = 200
 var num4 float64 = 123.456
 var b2 bool = false

 str = strconv.FormatInt(int64(num3), 10)
 fmt.Printf("%T %q\n", str, str)

 str = strconv.FormatFloat(num4, 'f', 10, 64)
 fmt.Printf("%T %q\n", str, str)

 str = strconv.FormatBool(b2)
 fmt.Printf("%T %q\n", str, str)
}
```

* **string 转成基本数据类型**
  * 使用 strconv 包的函数

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		string 类型转基本数据类型
		使用 strconv 包的函数
	*/
	var strToI = "100"
	var num int64
	var num1 int8

	// 转成整数
	num, _ = strconv.ParseInt(strToI, 10, 64)
	num1 = int8(num)
	fmt.Printf("%T %v\n", num1, num1)

	// 转成小数
	var strToF = "123.456"
	var f float64
	f, _ = strconv.ParseFloat(strToF, 64)
	fmt.Printf("%T %v\n", f, f)

	// 转成 bool
	var strToB = "true"
	var b bool
	b, _ = strconv.ParseBool(strToB)
	fmt.Printf("%T %v\n", b, b)
}
```

* **string 转基本数据类型的注意事项**

  * 在 将 string 类型转成基本数据类型时，要确保 string 类型能够转成有效的数据，比如 把 “123”，转成一个整数，但是不能把“dwj”转成一个整数，如果这样做，go 直接将其转成 0，其他类型也类似：float -> 0, bool -> false



### 指针

* 基本数据类型，变量存的就是值，也叫值类型
* 获取变量的地址，用 & 比如：var num int 获取 num 的地址：&num
* 指针类型，变量存的是一个地址，这个地址指向的空间存的才是值，比如：var ptr *int = &num
* 获取指针类型所指向的值，使用 *, 比如：var ptr *int, 使用 *ptr 获取 ptr 存储的内容(地址) 指向的值
* & 取地址， * 取值
* 值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应的指针就是 *int，float32 对应的指针类型就是 *float...
* 值类型包括：基本数据类型 int，float，bool，string、数组和结构体struct

```go
package main

import "fmt"

func main() {
 /*
  基本数据类型，变量存的就是值，也叫值类型
  获取变量的地址，用 & 比如：var num int 获取 num 的地址：&num
  指针类型，变量存的是一个地址，这个地址指向的空间存的才是值，比如：var ptr *int = &num
  获取指针类型所指向的值，使用 *, 比如：var ptr *int, 使用 *ptr 获取 ptr 存储的内容(地址) 指向的值
  & 取地址 * 取值  *ptr 可以访问到那个空间
 */
 var num = 1

 // ptr 是个指针变量
 var ptr *int = &num
 fmt.Println(ptr) // 0x1400001c070

 fmt.Println(*ptr) // 1
  
 var num1 int8 = 100
 var ptr1 *int8 = &num1
 fmt.Println(ptr1, num1) // 0x1400001c078 100
 // 修改 num 的值
 *ptr1 = 110
 fmt.Println(num1) // 110
}
```

![pointer](https://raw.githubusercontent.com/dengwenj/Golang/main/src/go_code/project03/10pointer/pointerImage.png "pointerImage")


### 值类型和引用类型

* 值类型：基本数据类型 int，float，bool，string、数组和结构体struct
* 引用类型：指针，切片slice，map，管道chan，interface等都是引用类型

  **值类型和引用类型的特点**

* 值类型：变量直接存储值，内存通常在栈中分配
* 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆中分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，被垃圾回收器回收。



### 标识符

* go 对各种变量、方法、函数等命名时使用的字符序列称为标识符，凡是自己可以起名的地方都叫标识符

### 运算符

* 运算符是一种特殊的符合，用以表示数据的运算、赋值、比较
* 算符运算符、赋值运算符、比较运算符、逻辑运算符、位运算符、其他运算符

​	**算术运算符**

```go
package main

import "fmt"

func main() {
 // 如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
 fmt.Println(10 / 4) // 2
 var n1 float32 = 10 / 4
 fmt.Println(n1) // 2

 // 如果希望保留小数部分，则需要有浮点数参与运算
 var n2 float32 = 10 / 4.0
 fmt.Println(n2) // 2.5

 // a % b = a - a / b * b
 fmt.Println(10 % 3)   // 1
 fmt.Println(-10 % 3)  // -1
 fmt.Println(-10 % -3) // -1
 fmt.Println(10 % -3)  // 1
 
 // 在 go 中， ++ 和 -- 只能独立使用
 //var i = 10
 //var x int
 //x = i++ // 只能独立使用
 //i++
 //if i++ > 0 { // 只能独立使用
 //
 //}
}
```

**关系运算符（比较运算符）**

* 关系运算符的结果都是 bool 类型，关系表达式经常用在 if 结构中或 for循环结构中
* == != < > <= >= 结果都是布尔值  1 != 2

**逻辑运算符**

* 用于连接多个条件（一般来讲就是关系表达式），最终的结果也是一个 bool 值
* && 逻辑与、 || 逻辑或、 !逻辑非

* && 也叫 短路与：如果第一个条件为 false，则第二个条件不会判断，最终结果为 false
* || 也叫短路或：如果第一个条件为 true，则第二个条件不会判断，最终结果为 true

**赋值运算符**

* 赋值运算符就是将某个运算后的值，赋给指定的变量
* 左边只能是变量，右边可以是变量、表达式、常量值

**运算符优先级**

* 括号，++，--
* 单目运算
* 算术元运算符
* 移位运算
* 关系运算符
* 位运算符（&、|、 ^、 <<、 >>、涉及到二进制）
* 逻辑运算符
* 赋值运算符
* 逗号

**其他运算符**

* &地址符 返回变量存储地址，&a 将给出变量的实际地址
* *取值符 指针变量， *a是一个指针变量

**键盘输入语句**

```go
package main

import "fmt"

func main() {
 var name string
 var age byte
 var sal float32
 var isPass bool

 //fmt.Println("请输入姓名")
 //fmt.Scanln(&name)
 //
 //fmt.Println("请输入年龄")
 //fmt.Scanln(&age)
 //
 //fmt.Println("请输入薪水")
 //fmt.Scanln(&sal)
 //
 //fmt.Println("是否通过")
 //fmt.Scanln(&isPass)

 //fmt.Printf("姓名%v\n 年龄%v\n 薪水%v\n 是否通过%v\n", name, age, sal, isPass)
  
 fmt.Println("请输入姓名、年龄、薪水、是否通过")
 fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
 fmt.Printf("姓名%v\n 年龄%v\n 薪水%v\n 是否通过%v\n", name, age, sal, isPass)
}
```

**进制**

* 二进制：0，1，满2进1
* 十进制：0-9，满10进1
* 八进制：0-7，满8进1，以数字0开头表示
* 十六进制：0-9及A-F（不区分大小写），满16进1，以0x或0X开头表示

```
0xD24
4 0100
2 0010
13 1101
110100100100
```

**其他进制转十进制**

* 二进制：从最低位开始（右边的），将每个位上的数提取出来，乘以2的（位数-1）次方然后求和
* 八进制：从最低位开始（右边的），将每个位上的数提取出来，乘以8的（位数-1）次方然后求和
* 十六进制：从最低位开始（右边的），将每个位上的数提取出来，乘以16的（位数-1）次方然后求和

**十进制转其他进制**

* 二进制：将该数不断除以2，直到商为0为止，然后将每步得到的余数倒过来，就是对应的二进制
* 八进制：将该数不断除以8，直到商为0为止，然后将每步得到的余数倒过来，就是对应的二进制
* 十六进制：将该数不断除以16，直到商为0为止，然后将每步得到的余数倒过来，就是对应的二进制

**二进制转换成八进制、十六进制**

* 二进制转换成八进制：将二进制数每三位一组（从右开始组合），进行10进制的转换然后进行拼接。
* 二进制转换成十六进制：将二进制数每四位一组（从右开始组合），进行10进制的转换然后进行拼接。

**八进制、十六进制转换成二进制**

* 八进制转换成二进制：将八进制数每一位，转成对应的一个3位的二进制数即可。
* 十六进制转换成二进制：将十六进制数每一位，转成对应的一个4位的二进制数即可。



