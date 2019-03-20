package main

import (
	"fmt"
	"os"
)

func main() {
	key := "PWD"
	dir := os.Getenv(key)
	fmt.Println(dir)
}
