package main

import "fmt"

//  for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	for _, v := range sl {
		fmt.Println(v)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
