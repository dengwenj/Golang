package atomic

import "sync/atomic"

func LAtomicDetail() {
	var num int32 = 100

	//atomic.AddInt32(&num, 1)
	//atomic.AddInt32(&num, -1)
	//atomic.LoadInt32(&num) // 读
	//atomic.StoreInt32(&num, 200) // 存
	//比较交换
	atomic.CompareAndSwapInt32(&num, num, 300) // 300
	println(num)
}
