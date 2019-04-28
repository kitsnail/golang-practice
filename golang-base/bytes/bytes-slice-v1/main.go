package main

import (
	"fmt"
)

func main() {
	b1 := []byte{'a', 'b', 'c', 'd'}
	fmt.Println("byte slice: ", b1)

	for i, b := range b1 {
		fmt.Println(i, ":", b)
	}
}
