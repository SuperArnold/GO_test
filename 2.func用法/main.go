package main

import "fmt"

func main() {
	a := 1
	b := 2
	a, b = b, a //用這方法來執行
	fmt.Println(a)
	fmt.Println(b)
}
