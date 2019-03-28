package main

import "fmt"

func main() {
	bchan := make(chan int, 2)
	go func() {
		bchan <- 1
		bchan <- 2
		bchan <- 3
	}()

	fmt.Println(<-bchan)
	fmt.Println(<-bchan)
	fmt.Println(<-bchan)
}
