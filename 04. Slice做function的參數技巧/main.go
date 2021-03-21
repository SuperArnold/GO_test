package main

import (
	"fmt"
)

func add(foo []string) {
	foo = append(foo, "c")
	fmt.Println("Modify:", foo)
}
func main() {
	foo := []string{"a", "b"}
	fmt.Println("before foo :", foo)
	add(foo)
	fmt.Println("after foo :", foo)

	bar := foo[:1] //將foo第一個給bar
	fmt.Println("bar:", bar)
	s1 := append(bar, "c")
	fmt.Println("foo:", foo) //foo跟s1一樣的原因是bar跟s1都是用foo的址
	fmt.Println("s1:", s1)
	s2 := append(bar, "d") //同s1
	fmt.Println("foo:", foo)
	fmt.Println("s2:", s2)
	s3 := append(bar, "e", "f")
	fmt.Println("foo:", foo) //為什麼跟s3不一樣了？
	fmt.Println("s3:", s3)
}
