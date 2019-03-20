package main

import (
	"fmt"
	"path"
)

func main() {
	p1 := "awll/awdjkk/a.go"

	fmt.Println("path:", p1)
	fmt.Println("path base:", path.Base(p1))
	fmt.Println("path is ABS:", path.IsAbs(p1))
	fmt.Println("path dir:", path.Dir(p1))
	fmt.Println("path ext:", path.Ext(p1))
}
