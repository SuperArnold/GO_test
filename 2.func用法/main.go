package main

import (
	"fmt"
)

func main() {
	//一般範例
	bar := func() {
		fmt.Println("Hello")
	}
	bar()

	//不用參數接func卻可以執行的方法
	func() {
		fmt.Println("World")
	}() //這裡一定要加

	//使用go routine的方式在背景執行
	go func(i, j int) int {
		return i + j
	}(2, 2) //這裡一定要加
}
