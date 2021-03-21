package main

import (
	"fmt"
)

func add(i int) []int {
	// chan := make(chan int, i)
	var ints []int
	channel := make(chan int, i)

	for n := 0; n < i; n++ {
		go func(channel chan<- int, j int) {
			channel <- j
		}(channel, n)
	}

	for m := range channel {
		ints = append(ints, m)

		if len(ints) == i {
			break
		}
	}
	close(channel)
	return ints
}

func main() {
	foo := add(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}
