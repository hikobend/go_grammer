// ユーザーから入力を受け取る

package main

// fmt : 文法を呼び出す。
// math/rand : 乱数を生成.乱数を生成するために使うシード値がデフォルトで固定されている。
// time : 時間

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1
	count := 0

	// 繰り返し行う
	for {
		var guess int

		// Scanf : 文字列を入力させる
		fmt.Print("数字を入力してください")
		fmt.Scanf("%v", &guess)
		count++
		// 条件分岐
		if answer == guess {
			fmt.Printf(("正解！！！ %v回で正解しました!!\n"), count)
			break
		} else if answer > guess {
			fmt.Println(("正解は大きい値のようです"))
		} else {
			fmt.Println(("正解は 小さな値のようです"))
		}
	}

}
