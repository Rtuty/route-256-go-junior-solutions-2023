package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)
	go testChank(ch)

	for i := range ch {
		fmt.Printf("Chank data: %d \n", i)
	}
}

func testChank(ch chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
	close(ch)
}
