// 関数

package main

import (
	"fmt"
)

// 自分で作成
func hi() {
	fmt.Println("hi")
}

// オプションを追加

func hello(name string) {
	fmt.Println("hi!" + name)
}

func test(testname string, content string) {
	fmt.Println(testname + "の名前は" + content + "です")
}

// 返り値
// 返り値の方を関数名の後に設定する必要がある
func food(like string) string {
	msg := "好きな食べものは" + like + "です"
	return msg
}

// 名前付きリターン
func sports(like string) (msg2 string) {
	msg2 = "好きなスポーツは" + like + "です"
	return msg2
}

func main() {
	hi()
	hello("kohei")
	test("アイコンのキャラ", "dat")
	fmt.Println(food("たまご"))
	fmt.Println(sports("野球"))
}
