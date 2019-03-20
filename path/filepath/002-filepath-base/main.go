package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//p := "/data/worklib/src/practices/golang/stdlib/path/filepath/001-filepath-base/a"
	p := "aaaa"
	pfix := "/data"

	fmt.Println("file path:", p)
	abs, err := filepath.Abs(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file path ABS:", abs)
	fmt.Println("file path base:", filepath.Base(p))
	fmt.Println("file path dir:", filepath.Dir(p))
	link, err := filepath.EvalSymlinks(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file path eval symlinks:", link)
	fmt.Println("file path exttend:", filepath.Ext(p))
	fmt.Println("file path from slash:", filepath.FromSlash(p))
	glob, err := filepath.Glob(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file path Glob:", glob)
	fmt.Println("file path has prefix:", filepath.HasPrefix(p, pfix))
	fmt.Println("file path is abs:", filepath.IsAbs(p))
	dir, file := filepath.Split(p)
	fmt.Println("file path split dir:", dir)
	fmt.Println("file path split file:", file)
	fmt.Println("file path split list:", filepath.SplitList(p))
	fmt.Println("file path to slash:", filepath.ToSlash(p))
	fmt.Println("file path volumename:", filepath.VolumeName(p))
}
