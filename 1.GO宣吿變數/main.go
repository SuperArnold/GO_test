package main

import "fmt"

const (
	monday = iota + 1
	tusday //不需再打
)

func main() {
	foo := "Arnold"
	bar := 100
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(monday)
	fmt.Println(tusday)
	fmt.Println("Done")
}
