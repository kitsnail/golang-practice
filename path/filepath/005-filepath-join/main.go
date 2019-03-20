package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	rootpath := `E:\apps\film`
	name := "askdwkd.ask"
	fmt.Println(filepath.ToSlash(rootpath + filepath.Separator + name))
}
