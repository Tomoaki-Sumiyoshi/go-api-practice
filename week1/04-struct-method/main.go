package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	user := User{FirstName: "Taro", LastName: "Yamada"}

	fmt.Println(user.FullName())
}
