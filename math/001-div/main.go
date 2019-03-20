package main

import "fmt"

func main() {
	size := 2682517688
	thread := 25

	result := size / thread

	fmt.Printf("%d / %d = %d", size, thread, result)
}
