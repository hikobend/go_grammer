// 並列処理
// goroutine

package main

import (
	"fmt"
	"time"
)

func task1() {
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished")
}

func task2() {
	fmt.Println("task2 finished")
}

func main() {
	go task1()
	go task2()

	// goroutineが終わる前に終わらないようにする
	time.Sleep(time.Second * 3)
}
