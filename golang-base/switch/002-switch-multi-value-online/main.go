package main

import (
	"fmt"
)

func main() {
	slices := []int{10, 9, 3, 4, 8, 0}

	for _, slice := range slices {
		fmt.Println("--->", slice)
		switch slice {
		case 1, 2, 3, 4:
			fmt.Println(slice, "the first group")
		case 5, 6, 7, 8:
			fmt.Println(slice, "the second group")
		default:
			if slice > 9 {
				fmt.Println(slice, "the third group")
			}
		}
	}
}
