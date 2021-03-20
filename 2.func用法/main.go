package main

import "fmt"

func main() {
	bar := func(i, j int) int { //執行將func寫進去
		return i + j
	}
	fmt.Printf("%T\n", bar)
	fmt.Println(bar(1, 5))
}
