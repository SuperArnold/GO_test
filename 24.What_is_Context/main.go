package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("stop job.")
				return
			default:
				fmt.Println("still work")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	stop <- true
	time.Sleep(1 * time.Second)
	fmt.Println("All Job Done.")
}
