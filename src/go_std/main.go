package main

import (
	"fmt"
)

func 讲了一个故事(num int) string {
	if num == 1 {
		return "师父，你已经讲了一万遍啦！"
	}
	return `
		从前有座山
		山上有座庙
		庙里有个老和尚和小和尚
		老和尚在给小和尚讲故事🫡
	` + 讲了一个故事(num-1)
}

func main() {
	fmt.Printf("%v\n", 讲了一个故事(10001))
}
