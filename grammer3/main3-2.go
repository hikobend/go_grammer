// マップ
package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["kohei"] = 24
	m["study-day"] = 2

	// キーを指定しながら宣言
	n := map[string]int{"dat-age": 100, "dat-birth": 927}

	// キーの存在を調べる
	p, ok := m["kohei"]
	q, ok := m["dat-age"]

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(n)

	delete(m, "study-day")

	fmt.Println(m)

	fmt.Println(p)
	fmt.Println(ok)

	fmt.Println(q)
	fmt.Println(ok)

}
