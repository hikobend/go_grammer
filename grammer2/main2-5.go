package main

import (
	"fmt"
)

// 2つの数値をひっくり返す
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(swap(2, 5))

	f := func(x, y int) (int, int) {
		return y, x
	}
	fmt.Println((f(3, 9)))

	// 即時関数
	func(msg string) {
		fmt.Println(msg)
	}("kohei")
}
