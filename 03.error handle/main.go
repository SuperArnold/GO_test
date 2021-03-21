package main

import (
	"errors"
	"fmt"
)

type errUserExist struct {
	Name string
}

func (e errUserExist) Error() string {
	return fmt.Sprintf("Username %s is exist.", e.Name)
}
func isErrUserExist(err error) bool { //判斷是不是自己定義的，error這個interface本身就有判斷
	_, ok := err.(errUserExist)
	return ok
}
func checkUserExist(username string) (bool, error) {
	if username == "foo" {
		return true, errUserExist{Name: username}
	}
	if username == "foo" {
		return true, errors.New("Username bar is exist.")
	}
	return false, nil
}
func main() {
	if _, foo_err := checkUserExist("foo"); foo_err != nil {
		if isErrUserExist(foo_err) {
			fmt.Println(foo_err)
		}
	}
	if _, bar_err := checkUserExist("bar"); bar_err != nil {
		if isErrUserExist(bar_err) {
			fmt.Println(bar_err)
		}
	} else {
		fmt.Println("bar is not struct error")
	}
}
