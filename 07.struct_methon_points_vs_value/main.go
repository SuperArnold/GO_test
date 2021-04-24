package main

import (
	"fmt"
	"time"
)

type email struct {
	from string
	to   string
}

func (e email) From(s string) email {
	e.from = s
	return e
}
func (e email) To(s string) email {
	e.to = s
	return e
}
func (e email) Send() {
	fmt.Printf("Send emaol From %s To %s\n", e.from, e.to)
}
func main() {
	email := &email{}
	for i := 1; i <= 10; i++ {
		go func(i int) {
			email.From(fmt.Sprintf("User%02d", i)).
				To(fmt.Sprintf("User%02d", i+1)).
				Send()
		}(i)
	}
	time.Sleep(1 * time.Second)
}
