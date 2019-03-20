package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := "/data/worklib/src/practices/golang/stdlib/path/filepath/001-filepath-base/main.go"
	fmt.Println(p)

	fn := filepath.Base(p)

	fmt.Println(fn)
}
