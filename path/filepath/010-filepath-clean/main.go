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
		"/home/wanghui/../..",
		"/home/wanghui/../../..",
		"/home/wanghui/../../../..",
		"/home/wanghui/.",
		"/home/wanghui/./",
		"wanghui",
		"/home/wanghui/../"}

	for n, p := range paths {
		fmt.Printf("%d path: %s\n", n, p)
		fmt.Printf("%d path clean: %s\n", n, filepath.Clean(p))
	}
}
