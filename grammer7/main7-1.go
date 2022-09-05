package main

// if

import "fmt"

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I dont know")
	}

	b := 100
	if b == 100 {
		fmt.Println("100です")
	}
}
