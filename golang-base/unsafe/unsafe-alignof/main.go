package main

import (
	"fmt"
	"unsafe"
)

type part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type part2 struct {
	e byte
	c int8
	a bool
	b int32
	d int64
}

func main() {
	p1 := part1{}
	p2 := part2{}

	fmt.Printf("part1 size: %d,aligin: %d\n", unsafe.Sizeof(p1), unsafe.Alignof(p1))
	fmt.Printf("part2 size: %d,aligin: %d\n", unsafe.Sizeof(p2), unsafe.Alignof(p2))
}
