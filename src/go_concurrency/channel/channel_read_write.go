package channel

func LChannelReadWrite() {
	// 通道遍历

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 写入通道
		}
		// 通道写完了过后要关闭 close()
		close(c)
	}()

	//	读取通道里面的值
	//cData := <-c
	//fmt.Printf("cData%v\n", cData)
	//遍历读取
	for i := range c {
		num := i
		println(num)
	}

	//for {
	//	data, ok := <-c
	//	if ok {
	//		fmt.Printf("%v\n", data)
	//	} else {
	//		break
	//	}
	//}
}
