package main

import "fmt"

// Generics

/*
func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func PrintSliceString(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
*/

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	// PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSliceString([]string{"a", "b", "c"})

	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice([]string{"a", "b", "c"})
}
