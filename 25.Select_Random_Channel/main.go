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
		case <-tick.C:
			fmt.Printf("%d is tick\n", i)
		}

		time.Sleep(200 * time.Millisecond)
	}
	close(channel)
	tick.Stop()
}
