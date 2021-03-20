package main

import (
	"fmt"
	"strings"
)

func getUserList(username, email string) string {
	sql := "select * from"
	where := []string{}
	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}
	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}
	return sql + " where " + strings.Join(where, " or ")
}
func main() {
	fmt.Println(getUserList("Arnold", ""))
	fmt.Println(getUserList("Arnold", "arnold@gmail.com"))
}
