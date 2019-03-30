package main

import (
	"fmt"
	"time"
)

func main() {
	bchan := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		go printer(i, bchan)
	}
	time.Sleep(1 * time.Hour)
}

func printer(x int, token chan struct{}) {
	token <- struct{}{}
	fmt.Printf("%d -> squares: %d\n", x, x*x)
	time.Sleep(1 * time.Second)
	<-token
}
