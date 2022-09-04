// range

package main

import (
	"fmt"
)

func main() {
	s := []int{4, 7, 2, 9}

	for i, v := range s {
		fmt.Println(i, v)
	}

	for _, v := range s {
		fmt.Println(v)
	}

	m := map[string]int{"kohei_age": 24, "study_day": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
