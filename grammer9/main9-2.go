package main

import "fmt"

// 構造体

type User struct {
	Name string
	Age  int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)

	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{
		Name: "user2",
		Age:  15,
	}
	fmt.Println(user2)

	user3 := User{
		"user3",
		49,
	}
	fmt.Println(user3)

	user4 := User{
		Name: "user3",
	}
	fmt.Println(user4)

	user5 := new(User)
	fmt.Println(user5)

	user6 := &User{}
	fmt.Println(user6)

	UpdateUser(user1)
	UpdateUser2(user6)

	fmt.Println(user1)
	fmt.Println(user6)
}
