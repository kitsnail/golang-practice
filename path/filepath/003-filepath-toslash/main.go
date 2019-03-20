package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//s := `http://www.site.com/a/b/c/d`
	s := `D:\film\winmeda\aasdwd.flv`
	// 下面这句用于 Windows 系统
	s = filepath.FromSlash(s)
	fmt.Println(s) // /a/b/c/d 或 \a\b\c\d
	s = filepath.Dir(s)
	fmt.Println(s) // /a/b/c/d
	s = filepath.ToSlash(s)
	fmt.Println(s) // /a/b/c/d
}
