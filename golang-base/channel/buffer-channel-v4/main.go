package main

import (
	"fmt"
	"time"
)

func main() {
	bchan := make(chan int, 2)
	go createWorkers(bchan)
	for v := range bchan {
		fmt.Println("==>read value", v, "from ch")
	}
}

func createWorkers(ch chan int) {
	finish := make(chan bool, 2)
	for i := 0; i < 5; i++ {
		select {
		case <-finish:
			ch <- i
		default:
			go worker(i, finish)
		}
	}
	close(ch)
}

func worker(id int, fc chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("successfully wrote", id, "to ch")
	fc <- true
}
