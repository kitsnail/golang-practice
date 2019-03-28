package main

import "fmt"

func main() {
	bchan := make(chan int, 2)

	for i := 0; i < 10; i++ {
		go func(n int) {
			bchan <- n
		}(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-bchan)
	}
	close(bchan)
}
