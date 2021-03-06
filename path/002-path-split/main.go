package main

import (
	"fmt"
	"path"
)

func main() {
	paths := []string{"/home/wanghui",
		"/home/wanghui/",
		"../home/wanghui",
		"../home/wanghui/",
		"home/wanghui",
		"home/wanghui/",
		"wanghui"
		"/home/wanghui/..",
		"/home/wanghui/../"}

	for n, p := range paths {
		fmt.Printf("%d path: %s\n", n, p)
		dir, file := path.Split(p)
		fmt.Printf("%d path split, dir: %s,files:%s\n", n, dir, file)
	}
}
