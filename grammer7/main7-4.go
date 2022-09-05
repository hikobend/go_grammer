package main

import "fmt"

// 式スイッチ

func main() {
	n := 4

	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	}

	switch x := 2; x {
	case 1, 2:
		fmt.Println("1 ro 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	}

	fmt.Println(n)

	// fmt.Println(x)

	a := 6
	switch {
	case a > 0 && a < 4:
		fmt.Println("0 < a < 4")
	case a > 3 && a < 8:
		fmt.Println("3 < a < 8")
	}

	switch {
	case 1, 2:
		fmt.Println("1 ro 2")
	case a > 0 && a < 4:
		fmt.Println("0 < a < 4")
	case a > 3 && a < 8:
		fmt.Println("3 < a < 8")
	}
}
