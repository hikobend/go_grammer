package main

// fmt : 変数宣言

import (
	"fmt"
)

func main() {
	// 変数
	a := 10
	b := 12.3
	c := "hello world"
	var d bool

	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)

	// 定数
	const name = "kohei"
	fmt.Println(name)

	const (
		k = iota
		o
		h
	)

	fmt.Printf("k: %d, o:%d, h:%d\n", k, o, h)
}
