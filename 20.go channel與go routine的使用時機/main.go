package main

import (
	"fmt"
	"sync"
)

func add(i int) []int {
	// chan := make(chan int, i)
	var ints []int
	var wg sync.WaitGroup
	wg.Add(i)
	for n := 0; n < i; n++ {
		go func(j int) {
			defer wg.Done()
			ints = append(ints, j)
		}(n)
	}
	wg.Wait()
	return ints
}

func main() {
	foo := add(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}
