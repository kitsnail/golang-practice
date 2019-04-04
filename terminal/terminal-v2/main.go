package main

import (
	"bytes"
	"os"
	"time"
)

const (
	eraseLine         = "\x1b[2K"
	moveToStartOfLine = "\x1b[0G"
	moveUp            = "\x1b[1A"
)

func main() {
	var buf bytes.Buffer
	var buf2 bytes.Buffer
	buf.WriteString("aaaa")
	buf.WriteString("\n")
	os.Stdout.Write(buf.Bytes())
	time.Sleep(3 * time.Second)
	buf2.WriteString(moveUp)
	os.Stdout.Write(buf2.Bytes())
	//fmt.Println("\033[2J\033[6;3HHello")

}
