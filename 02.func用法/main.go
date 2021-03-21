package main

import (
	"fmt"
)

type errUserExist struct {
	Name string
}

func (e errUserExist) Error() string {
	return fmt.Sprintf("Username %s is exist.", e.Name)
}
func checkUserExist(username string) (bool, error) {
	if username == "foo" {
		return true, errUserExist{Name: username}
	}
	return false, nil
}
func main() {
	if _, foo_err := checkUserExist("foo"); foo_err != nil {
		fmt.Println(foo_err)
	}
}
