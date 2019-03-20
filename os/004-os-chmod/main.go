package main

import (
	"fmt"
	"os"
)

var (
	PubKeyPerm os.FileMode = 0600
)

func main() {
	file1 := "file.data"

	fi, err := os.Stat(file1)
	if err != nil {
		panic(err)
	}
	// show default perm
	fmt.Println("show default perm:", fi.Mode())

	if fi.Mode().Perm() != PubKeyPerm {
		fmt.Println("the defult perm is error")
	}

	// change file perm to 0600
	// '-rw-------'

	//err = os.Chmod(file1, 0600)
	err = os.Chmod(file1, PubKeyPerm)
	if err != nil {
		panic(err)
	}

	fi2, err := os.Stat(file1)
	if err != nil {
		panic(err)
	}

	// show changed file perm
	fmt.Println("show fi1 perm status:", fi.Mode())

	// show fi2 perm status
	fmt.Println("show fi2 perm status:", fi2.Mode())
}
