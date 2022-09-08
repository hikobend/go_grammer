package main

import (
	"fmt"

	"example.com/m/grammer14/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()

	/*
		u, _ := models.GetUser(1)

		fmt.Println(u)

		u.Name = "Test2"
		u.Email = "test@test.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*

		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	/*

		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")

	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}
}
