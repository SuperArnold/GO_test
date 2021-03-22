package main

import (
	"fmt"
)

func init() {
	global = 0
}

var global = convert()

func convert() int {
	return 100
}
func main() {
	fmt.Println("Global is ", global)
}
