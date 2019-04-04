package main

import (
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Println("Print hello world 1")
	ln := tm.CurrentHeight()
	tm.Flush()
	time.Sleep(time.Second * 1)

	tm.MoveCursorUp(ln)
	tm.Println("Print hello world 2")
	tm.Flush()
}
