package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{"/home/wanghui",
		"/home/wanghui/",
		"../home/wanghui",
		"../home/wanghui/",
		"home/wanghui",
		"home/wanghui/",
		"/home/wanghui/..",
		"wanghui",
		"/home/wanghui/../"}

	for n, p := range paths {
		fmt.Printf("%d path: %s\n", n, p)
		list := filepath.SplitList(p)
		fmt.Printf("%d path splitList: %v\n", n, list)
	}
}
