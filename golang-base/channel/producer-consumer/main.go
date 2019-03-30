package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("write:", i)
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for value := range in {
		fmt.Println("read:", value)
	}
}
