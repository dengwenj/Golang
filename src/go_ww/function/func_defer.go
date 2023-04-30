package function

func LFuncDefer() {
	print("start...")
	defer print("我是 defer1")
	defer print("我是 defer2")
	defer print("我是 defer3")
	print("end...")
}
