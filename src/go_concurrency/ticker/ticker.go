package ticker

import (
	"fmt"
	"time"
)

func LTicker() {
	//ticker := time.NewTicker(time.Second)
	//
	//sum := 0
	//for {
	//	t := <-ticker.C
	//	fmt.Printf("%v\n", t)
	//	sum++
	//	if sum >= 3 {
	//		break
	//	}
	//}

	ticker := time.NewTicker(time.Second)
	chanInt := make(chan int)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for i := range chanInt {
		fmt.Printf("%v\n", i)
		sum += i
		if sum >= 10 {
			break
		}
	}
}
