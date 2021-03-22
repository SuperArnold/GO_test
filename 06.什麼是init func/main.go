package main

import (
	"fmt"
)

func init() {
	fmt.Println("init 1")
}
func init() {
	fmt.Println("init 2")
}
func main() {
	fmt.Println("main function")
}
