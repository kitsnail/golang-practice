package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fn := os.Args[1]
	fi, err := os.Stat(fn)
	if err != nil {
		log.Fatal(err)
	}
	m := fi.Mode()

	if m.IsRegular() {
		fmt.Println("regular file")
	}
}
