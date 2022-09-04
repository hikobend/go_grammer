// for (while文はない)

package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}
}
