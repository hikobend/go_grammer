// 変数

package main

import "fmt"

// 関数外では、型推論はできない。
// i5 = 500

var i5 int = 500

func outer() {
	var s4 string = "hello"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t)
	fmt.Println(f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
}
