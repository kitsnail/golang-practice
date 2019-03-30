package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	bchan := make(chan struct{}, 10)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go printer(&wg, i, bchan)
	}
	wg.Wait()
}

func printer(wg *sync.WaitGroup, x int, token chan struct{}) {
	defer wg.Done()
	token <- struct{}{}
	fmt.Printf("%d -> squares: %d\n", x, x*x)
	time.Sleep(1 * time.Second)
	<-token
}
