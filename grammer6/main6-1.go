package main

import "fmt"

// 関数

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	// returnの中身は省略できる
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(10, 3)
	fmt.Println(i2, i3)

	// 使わない文字は_
	i4, _ := Div(20, 3)
	fmt.Println(i4)

	i5 := Double(2000)
	fmt.Println(i5)

	Noreturn()
}
