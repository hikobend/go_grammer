// 型変換

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)

	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	j, _ := strconv.Atoi(s)
	fmt.Println(j)
	fmt.Printf("j = %T\n", j)

	var j2 int = 200
	s2 := strconv.Itoa(j2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "hello world"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}