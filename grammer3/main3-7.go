// 構造体

package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	// ポインタが帰ってくる
	u := new(user)

	// ポインタが蹴ってこない
	u.name = "kohei"

	// ポインタが蹴ってこない
	u.age = 24

	v := user{"dat", 100}

	w := user{name: "friend", age: 99}

	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)
}
