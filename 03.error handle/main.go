package main

import (
	"errors"
	"fmt"
)

func checkUserExist(username string) (bool, error) {
	if username == "foo" {
		return true, fmt.Errorf("Username %s is exist.", username) //使用fmt.errorf
	}
	if username == "bar" {
		return true, errors.New(fmt.Sprintf("Username %s is exist.", username)) //使用errors.New定義新的error
	}
	return false, nil
}
func main() {
	if _, foo_err := checkUserExist("foo"); foo_err != nil {
		fmt.Println(foo_err)
	}
	if _, bar_err := checkUserExist("bar"); bar_err != nil {
		fmt.Println(bar_err)
	}
}
