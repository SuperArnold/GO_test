package main

import "fmt"

func swap(i, j int) (int, int) {
	return j, i
}
func main() {
	a, b := swap(1, 2)
	fmt.Println(a)
	fmt.Println(b)
}
