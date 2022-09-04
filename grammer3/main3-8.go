// メソッド

package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u user) show() {
	fmt.Printf("name:%s, age:%d\n", u.name, u.age)
}

// 値渡し
func (u user) hit1() {
	u.age++
}

// 参照渡し
func (u *user) hit2() {
	u.age++
}

func main() {
	u := user{name: "kohei", age: 24}
	// show hitの順番によって結果が変わる。

	// u.hit2()
	u.hit1()
	u.show()
}
