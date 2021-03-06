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
		go squarer(&wg, i, bchan)
	}
	wg.Wait()
}

func squarer(wg *sync.WaitGroup, x int, token chan struct{}) {
	token <- struct{}{}
	defer func() { <-token }()
	defer wg.Done()

	fmt.Printf("%d -> squares: %d\n", x, x*x)
	time.Sleep(1 * time.Second)
}
