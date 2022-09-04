// if文
/*
> >=
< <=
==
!=
&&
||
*/

package main

import (
	"fmt"
)

func main() {

	score := 851

	if score > 720 {
		fmt.Println("合格です")
	} else if score > 650 {
		fmt.Println("残念...あともう少しでした...")
	} else {
		fmt.Println("出直してこいやぁぁぁ---!!!")
	}

	//  値をif文の中に入れる
	// if文の中でしか定義されていないことに注意する

	if score2 := 851; score2 > 750 {
		fmt.Println("合格!!!!!!!!!!です笑")
	} else if score2 > 620 {
		fmt.Println("残念.................あともう少しでした...")
	} else {
		fmt.Println("出直してこいやぁぁぁ---------------------!!!")
	}

}
