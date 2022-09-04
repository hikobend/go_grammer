// 配列
// スライス
//  make
// append
// copy

package main

import (
	"fmt"
)

func main() {
	// 別々に設定
	var a [5]int
	a[2] = 5
	a[4] = 18

	b := [...]int{1, 4, 2, 6, 5, 1, 8}
	s := b[2:4]
	t := b[:4]
	u := b[:]

	// 配列からスライスを切り出す
	v := make([]int, 3)
	w := []int{1, 3, 5}

	// append
	// 要素の末尾に要素を追加する
	// 同じ文字に追加する
	v = append(v, 14, 747, 53)

	// copy
	x := make([]int, len(v))
	y := copy(x, v)

	// スライスで値を変更すると元の値も変わる
	// s[1] = 100とすると, bも100になる
	s[1] = 100

	fmt.Println(a)
	fmt.Println(a[2])

	fmt.Println(b)
	// bの長さ
	fmt.Println(len(b))

	fmt.Println(s)
	// sの切り出しから最大に取り出せる値の個数
	fmt.Println(cap(s))

	fmt.Println(t)
	fmt.Println(u)

	fmt.Println(v)
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)

}
