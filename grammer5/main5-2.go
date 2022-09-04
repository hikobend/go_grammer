package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println(9 + 4)

	fmt.Println(9 - 4)

	fmt.Println(9 * 4)

	fmt.Println(9 / 4)

	fmt.Println(9 % 4)

	x := 0
	x += 5
	fmt.Println(x)

	x++
	fmt.Println(x)

	x--
	fmt.Println(x)

	y := "ABC"
	y += "DEF"
	fmt.Println(y)
}
