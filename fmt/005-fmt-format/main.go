package main

import "fmt"

func main() {
	slice := []string{"1", "11", "111", "1111", "11111", "111111"}

	for i, s := range slice {
		fmt.Printf("%3d %6s\n", i, s)
	}

	for i, s := range slice {
		fmt.Printf("%3d %s\n", i, s)
	}

	for i := 1; i < 10; i++ {
		fmt.Printf("%5d %5d\n", i, i)
	}
}
