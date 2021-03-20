package main

import "fmt"

func foo() func() int { //回傳func
	return func() int {
		return 100
	}
}
func main() {
	bar := foo()
	fmt.Printf("%T\n", bar) //印出func定義
	fmt.Println(bar())
}
