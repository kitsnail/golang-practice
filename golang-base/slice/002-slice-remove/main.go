package main

import (
	"fmt"
)

func main() {
	all := []string{}

	for i := 0; i < 1000; i++ {
		elem := fmt.Sprintf("node%d", i)
		all = append(all, elem)
	}

	filter := []string{"node300", "node400", "node600", "node800", "node900"}
	fmt.Println(all)
	ns := strSliceRemove(all, filter...)
	fmt.Println(ns)
}

func strSliceRemove(slice []string, elems ...string) []string {
	isInElems := make(map[string]bool)
	for _, elem := range elems {
		isInElems[elem] = true
	}

	w := 0
	for _, elem := range slice {
		if !isInElems[elem] {
			slice[w] = elem
			w += 1
		}
	}
	return slice[:w]
}
