// switch

package main

import (
	"fmt"
)

func main() {
	baseball := "O"
	switch baseball {
	case "s":
		fmt.Println("ストライク")
	case "b":
		fmt.Println("ボール")
	case "o", "O":
		fmt.Println("アウト")
	default:
		fmt.Println("？？？？？")
	}

	baseball_score := 1
	switch {
	case baseball_score > 4:
		fmt.Println("試合は乱打戦です")
	default:
		fmt.Println("試合は投手戦です")
	}
}
