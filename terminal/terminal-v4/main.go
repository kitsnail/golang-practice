package main

import (
	tm "github.com/buger/goterm"
)

func main() {
	tm.Println("The 1st line.")
	tm.Println("line number:", tm.CurrentHeight())
	tm.Println("The 2nd line.")
	tm.Println("line number:", tm.CurrentHeight())
	tm.Println("The 3rd line.")
	tm.Println("line number:", tm.CurrentHeight())
	tm.Flush()
}
