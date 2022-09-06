package main

import "fmt"

// 構造体埋め込み

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	t.User.SetName()
	fmt.Println(t)
}
