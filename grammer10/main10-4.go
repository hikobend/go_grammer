package main

import (
	"fmt"
	f "fmt"

	"example.com/m/grammer10/foo.go"
)

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	return b
}

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func main() {
	f.Println(foo.Max)
	// fmt.Println(foo.min)

	f.Println(foo.ReturnMin())
	fmt.Println(foo.Max)

	fmt.Println(appName())
	// fmt.Println(AppName, Version)

	fmt.Println("AAA")
}
