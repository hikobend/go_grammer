// 並列処理
// goroutine
// チャネル

package main

import (
	"fmt"
	"time"
)

func task1(result chan string) {
	time.Sleep(time.Second * 2)
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished")
}

func main() {
	result := make(chan string)

	go task1(result)
	go task2()

	fmt.Println(<-result)
	// goroutineが終わる前に終わらないようにする
	time.Sleep(time.Second * 3)
}
