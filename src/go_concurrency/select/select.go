package _select

import (
	"fmt"
	"time"
)

func LSelect() {
	cI := make(chan int)
	cS := make(chan string)

	go func() {
		cI <- 1
		cS <- "邓文杰"

		defer close(cI)
		defer close(cS)
	}()

	for {
		select {
		case cIData := <-cI:
			fmt.Printf("%v\n", cIData)
		case cSData := <-cS:
			fmt.Printf("%v\n", cSData)
		default:
			fmt.Printf("%v\n", "default")
		}
		time.Sleep(time.Second)
	}
}
