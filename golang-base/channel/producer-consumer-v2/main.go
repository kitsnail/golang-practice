package main

import (
	"fmt"
	"time"
)

func main() {
	tokens := make(chan struct{}, 3)
	go producer(tokens)
	consumer(tokens)
	time.Sleep(1 * time.Hour)
}

func producer(tokens chan struct{}) {
	//var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//wg.Add(1)
		tokens <- struct{}{}
		go func(i int) {
			multiply(i, tokens)
		}(i)
	}
	close(tokens)
	//wg.Wait()
}

func multiply(i int, out chan<- struct{}) int {
	defer func() { out <- struct{}{} }()
	time.Sleep(1 * time.Second)
	fmt.Println("write:", i)
	return i * i

}

func consumer(in <-chan struct{}) {
	for value := range in {
		fmt.Println("read:", value)
	}
}
