package main

import "fmt"

// copy

func main() {
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000

	fmt.Println(sl)

	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2)

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 5)
	fmt.Println(sl4)

	n := copy(sl4, sl3)

	fmt.Println(n, sl4)
}
