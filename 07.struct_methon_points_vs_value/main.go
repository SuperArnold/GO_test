package main

import (
	"fmt"
	"time"
)

type email struct {
	from string
	to   string
}

func (e *email) From(s string) {
	e.from = s
}
func (e *email) To(s string) {
	e.to = s
}
func (e *email) Send() {
	fmt.Printf("Send emaol From %s To %s\n", e.from, e.to)
}
func main() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			email := &email{} //移到這裡
			email.From(fmt.Sprintf("User%02d", i))
			email.To(fmt.Sprintf("User%02d", i+1))
			email.Send()
		}(i)
	}
	time.Sleep(1 * time.Second)
}
