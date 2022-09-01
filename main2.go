// ユーザーから入力を受け取る

package main

import "fmt"

// fmt : 文法を呼び出す。

func main() {
	var guess int

	// Scanf : 文字列を入力させる
	fmt.Print("数字を入力してください")
	fmt.Scanf("%v", &guess)
	fmt.Printf("あなたの入力した値は %v です\n", guess)
}
