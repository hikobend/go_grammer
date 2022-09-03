// 四則演算

/*
数値 + - * / %
文字列 +
論理値 AND : && OR : || NOT !
*/

package main

import (
	"fmt"
)

func main() {
	var x int
	x = 10 % 3
	fmt.Println(x)

	var s string
	s = "hello" + "world"
	fmt.Println(s)

}
