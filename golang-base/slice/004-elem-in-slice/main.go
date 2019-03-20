package main

import "fmt"

func main() {
	sli := []string{"a", "b", "c"}
	elms := []string{"r", "c", "f", "1"}
	for _, elm := range elms {
		fmt.Println(elm)
		fmt.Println(isElemInSlice(sli, elm))
	}
}

func isElemInSlice(slice []string, elem string) bool {
	smap := make(map[string]bool)
	for _, s := range slice {
		smap[s] = true
	}
	return smap[elem]
}
