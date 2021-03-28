package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func(ch chan int) {
		for {
			if v, ok := <-ch; ok {
				fmt.Printf("val = %d \n", v)
			}
		}
	}(channel)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 30; i++ {
		select {
		case channel <- i:
		}

		select {
		case <-tick.C:
			fmt.Printf("%d is tick\n", i)
		default: //避免卡在tick中
		}

		time.Sleep(200 * time.Millisecond)
	}
	close(channel)
	tick.Stop()
}
