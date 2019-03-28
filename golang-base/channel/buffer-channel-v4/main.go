package main

import (
	"fmt"
	"time"
)

func main() {
	bchan := make(chan int, 2)
	go worker(bchan)
	time.Sleep(2 * time.Second)
	for v := range bchan {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func worker(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
