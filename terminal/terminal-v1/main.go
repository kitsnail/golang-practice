package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fd := os.Stdout.Fd()
	fmt.Println("os stdout fd:", fd)
	ifd := int(fd)
	if terminal.IsTerminal(int(ifd)) {
		fmt.Println("current in terminail")
	}

	w, h, err := terminal.GetSize(ifd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("width:", w, "height:", h)

	state1, err := terminal.GetState(ifd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(state1)

	state2, err := terminal.MakeRaw(ifd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(state2)

}
