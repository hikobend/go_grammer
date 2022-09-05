package main

import "fmt"

// 型スイッチ

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3

	i := x.(int)

	fmt.Println(i + 2)
	// fmt.Println(x + 2)

	f, isfloat64 := x.(float64)
	fmt.Println(f, isfloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	} else {
		fmt.Println("I dont kown")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("iI dont know")
	}
}
