package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	key := "HOME"
	dir := os.Getenv(key)
	fmt.Println(key, dir)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pwd", pwd)
}
