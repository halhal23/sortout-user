package main

import (
	"fmt"
	"sortout-user/domain"
)

func main() {
	fmt.Println("hello world sortout-user")
	u := domain.NewUserModel("hiroharu", "hh@exam.com")
	fmt.Printf("hello %s", u.Name)
}
