package main

import "fmt"

// panic recover
// 難しすぎる笑

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("start")
}
