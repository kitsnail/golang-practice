package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice := append(slice1, slice2...)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice:", slice)
}
