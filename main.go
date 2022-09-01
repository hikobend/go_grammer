package main

// fmt : 文法を呼び出す。

import (
	"fmt"
)

func main() {
	// 変数の設定
	say := "こんにちは"

	// 文字列の中に変数を埋め込む場合、Printf
	// 変数を埋め込みたい場所に%v.埋め込みたい変数を, sayのように記入
	// 改行 : \n
	fmt.Println("say world")
	fmt.Printf("%v world\n", say)
	fmt.Printf("%v k!\n", say)
}
