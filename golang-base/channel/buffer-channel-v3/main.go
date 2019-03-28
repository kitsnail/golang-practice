package main

import (
	"fmt"
	"time"
)

func main() {
	bchan := make(chan int, 2)
	for i := 0; i < 10; i++ {
		go worker(bchan, i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-bchan)
	}
}

func worker(ch chan int, n int) {
	time.Sleep(1 * time.Second)
	ch <- n
}
