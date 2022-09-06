package main

import (
	"fmt"
	"log"
	"os"
)

// os

func main() {
	// Exit

	// os.Exit(1)
	// fmt.Println("Start")

	// log,Fatal

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Args
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// on.Argsに要素を表示
	// fmt.Printf("length=%d\n", len(os.Args))
	// os.Argsの中身を全て表示
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// go build -o getargs main12-1.go
	// ./getargs 123 456 789

	// ファイル操作
	// Open
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer f.Close()

	// ファイル操作
	// Create
	// f, _ := os.Create("foo.txt")

	// f.Write([]byte("Hello\n"))

	// f.WriteAt([]byte("Golang"), 6)

	// f.Seek(0, os.SEEK_END)

	// f.WriteString("Yeah")

	// ファイル操作
	// read
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)
	fmt.Println(string(bs2))
}
