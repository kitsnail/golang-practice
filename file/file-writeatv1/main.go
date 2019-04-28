package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "text.txt"

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b1 := []byte("0123456789")
	n, err := f.WriteAt(b1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("b1 byte: ", b1)
	fmt.Println("b1 write at ", n)

	b2 := []byte("aaaaaaaaaa")
	n, err = f.WriteAt(b2, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("b2 byte: ", b2)
	fmt.Println("b2 write at ", n)

	b3 := []byte("bbbbbbbbbb")
	n, err = f.WriteAt(b3, 200)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("b3 byte: ", b3)
	fmt.Println("b3 write at ", n)
}
