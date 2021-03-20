package main

import (
	"fmt"
	"strings"
)

type searchOpts struct {
	username string
	email    string
}

func getUserList(opt searchOpts) string {
	sql := "select * from"
	where := []string{}
	if opt.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opt.username))
	}
	if opt.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opt.email))
	}
	return sql + " where " + strings.Join(where, " or ")
}
func main() {
	fmt.Println(getUserList(searchOpts{
		username: "Arnold",
	}))
	fmt.Println(getUserList(searchOpts{
		username: "Arnold",
		email:    "arnold@gmail.com",
	}))
}
