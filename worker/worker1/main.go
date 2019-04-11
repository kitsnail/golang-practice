package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // 執行完一次就 -1
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)          // 每次執行都 + 1
		go process(i, &wg) // wg 一定要用 pointer，否則每個 goroutine 都會有各自的 WaitGroup
	}
	wg.Wait() // 會 wait 到 0 才會繼續下一步
	fmt.Println("All go routines finished executing")
}
