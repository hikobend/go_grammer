package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	zero := 0.0

	pinf := 1.0 / zero
	fmt.Println(pinf)

	minf := -1.0 / zero
	fmt.Println(minf)

	nan := zero / zero
	fmt.Println(nan)
}
