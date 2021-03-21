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
}
