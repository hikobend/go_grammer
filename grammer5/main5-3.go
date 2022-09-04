package main

import "fmt"

// 比較演算子
// 論理演算子

func main() {
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(4 <= 8)
	fmt.Println(4 >= 8)
	fmt.Println(true != false)

	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)

}
