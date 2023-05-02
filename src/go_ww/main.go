package main

import _type "go_ww/type"

var dwj = initVar()

func init() {
	println("我是初始化函数 init 系统调用的")
}

func initVar() string {
	println("变量最先")
	return "邓文杰"
}

func main() {
	//name1 := user.Hello()
	//fmt.Printf("name1: %v\n", name1)

	// 变量 var
	//learnvar.LearnVar()

	//	常量 const
	//learnconst.LearnConst()

	//数据类型
	//learndatatype.LearnDataType()

	//	数字类型
	//number.LearnNumber()

	//字符串类型
	//string2.LearnString()

	//格式转换
	//format.Format()

	// if else
	//flow.IfFlow()

	// switch
	//flow.SwitchFlow()

	// for
	//flow.ForFlow()

	//for range
	//flow.ForRangeFlow()

	// break
	//flow.BreakFlow()

	//continue
	//flow.ContinueFlow()

	//goto
	//flow.GotoFlow()

	//数组
	//arr.Array()
	//arr.Element()

	//切片
	//slice.Slice()

	//初始化切片
	//slice.InitSlice()

	//切片增删改查
	//slice.SliceCRUD()

	//map 类型
	//_map.LMap()

	//函数
	//function.LFunction()
	//function.LFuncReturn()
	//function.LFuncParams()
	//function.LFuncType()
	//function.LfuncG()
	//function.LFuncN()
	//function.LFuncClosure()
	//function.LFuncRuc()
	//function.LFuncDefer()

	//指针
	//pointer.LPointer()
	//pointer.LPointerArr()

	//类型
	_type.LTypeDefine()
}
