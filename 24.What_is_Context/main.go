package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Job 1 finish.")
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Job 2 finish.")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All Job Done.")
}
