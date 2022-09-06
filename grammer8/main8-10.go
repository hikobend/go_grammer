package main

// 全くわからん!!

import (
	"fmt"
	"time"
)

// close

func reciver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	// close(ch1)

	// ch1 <- 1

	// fmt.Println(<-ch1)

	// ok : false -> 閉じられている
	// i, ok := <-ch1
	// fmt.Println(i, ok)

	go reciver("1.goroutine", ch1)
	go reciver("2.goroutine", ch1)
	go reciver("3.goroutine", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}
