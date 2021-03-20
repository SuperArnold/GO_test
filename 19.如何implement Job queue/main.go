package main

import (
	"fmt"
	"time"
)

func worker(jobchan_work <-chan int) {
	for i := range jobchan_work {
		fmt.Printf("Worker is %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
func enqueue(job int, jobchan_queue chan<- int) bool {
	select {
	case jobchan_queue <- job:
		return true
	default:
		return false
	}
}
func main() {
	jobchan := make(chan int, 1)
	go worker(jobchan)
	fmt.Println(enqueue(1, jobchan))
	fmt.Println(enqueue(2, jobchan))
	fmt.Println(enqueue(3, jobchan))
	time.Sleep(5 * time.Second)
}
