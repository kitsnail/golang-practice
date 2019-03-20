package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fp := `D:\apps\film\soe-32342.rmvb`
	fmt.Println(fp)
	fp = filepath.FromSlash(fp)
	fmt.Println(filepath.Ext(fp))
	fmt.Println(filepath.Base(fp))
}
