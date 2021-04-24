package main

import (
	"fmt"
)

type car struct {
	name string
}

func (c car) setCarName01(s string) { //傳value
	c.name = s
}
func (c *car) setCarName02(s string) { //傳pointer
	c.name = s
}
func main() {
	c := &car{}
	c.setCarName01("foo")
	fmt.Println(c.name)
	c.setCarName02("bar")
	fmt.Println(c.name)
}
