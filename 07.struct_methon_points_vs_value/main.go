package main

import (
	"fmt"
)

type car struct {
	name string
}

func (c car) setCarName01(s string) { //這裡其實是有兩個struct，而改的是這個function內的struct
	c.name = s
}
func (c *car) setCarName02(s string) {
	c.name = s
}
func main() {
	c := &car{}
	c.setCarName02("bar")
	fmt.Println(c.name)
	c.setCarName01("foo")
	fmt.Println(c.name)
}
