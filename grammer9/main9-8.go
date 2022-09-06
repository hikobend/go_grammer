package main

import "fmt"

// 構造体
// 独自性

type MyInt int

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// i := 100
	// fmt.Println(mi + i)
}
