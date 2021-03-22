package main

import (
	"fmt"
)

func check(s int) {
	switch s {
	case 0, 1:
		fmt.Println("Value is ", s)
	}
}
func main() {
	check(0)
	check(1)
}
