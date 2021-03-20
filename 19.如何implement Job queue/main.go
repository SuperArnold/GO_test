package main

import (
	"fmt"
	"time"
)

func worker(jobchan_work <-chan int) {
	for i := range jobchan_work {
		fmt.Printf("Worker is %d\n", i)
	}
}
func main() {
	jobchan := make(chan int, 1024)
	go worker(jobchan)
	jobchan <- 1
	fmt.Println("main is 1")
	jobchan <- 2
	fmt.Println("main is 2")
	time.Sleep(2 * time.Second)
}
