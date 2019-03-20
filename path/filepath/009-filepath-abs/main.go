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
		ap, err := filepath.Abs(p)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%d path ABS: %v\n", n, ap)
	}
}
