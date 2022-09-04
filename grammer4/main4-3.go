package main

import "fmt"

// 型

func main() {
	var i int = 100
	var i2 int64 = 200

	fmt.Println(i + 10000)
	// fmt.Println(i + i2

	// %Tは型を表示する
	fmt.Printf("%T\n", i2)

	// 型を変更する
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))
}
