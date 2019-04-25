package main

import "fmt"

type Interface1 interface {
	Method1()
}

type Interface2 interface {
	Method2()
}

type Struct struct{}

func (Struct) Method1() {
	fmt.Println("Struct Method1")
}

func (Struct) Method2() {
	fmt.Println("Struct Method2")
}

func Func1(i Interface1) {
	i.Method1()
}

func Func2(i Interface2) {
	i.Method2()
}

func main() {
	t := Struct{}
	Func1(t)
	Func2(t)
}
