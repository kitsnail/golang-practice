package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "text.txt"
	write(file)
	fmt.Println("------------------------")
	fstat(file)
}

func fstat(file string) {
	fi, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name:", fi.Name())
	fmt.Println("Size:", fi.Size())
	fmt.Println("Mode:", fi.Mode())
	fmt.Println("ModTime:", fi.ModTime())
	fmt.Println("isDir:", fi.IsDir())
	fmt.Println("Sys:", fi.Sys())
}

func write(file string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b1 := []byte("0000000000")
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

	b12 := []byte("1111111111")
	n, err = f.WriteAt(b12, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("b12 byte: ", b12)
	fmt.Println("b12 write at ", n)
}
