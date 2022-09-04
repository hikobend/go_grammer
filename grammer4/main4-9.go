package main

import "fmt"

// インターフェース
// 全ての型を汎用的に表す。

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

}
