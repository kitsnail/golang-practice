package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]

	fmt.Println("strings EqualFold:", strings.EqualFold(str1, str2))
}
