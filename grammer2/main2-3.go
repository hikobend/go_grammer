// ポインタ
// アドレスを指し示す変数

package main

import (
	"fmt"
)

func main() {
	a := 5
	var pa *int
	// &a は aのアドレス
	// paの領域にあるデータの値は*pa
	pa = &a

	fmt.Println(pa)
	fmt.Println(*pa)

}
