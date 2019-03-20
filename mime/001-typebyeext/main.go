package main

import (
	"fmt"
	"mime"
)

func main() {
	s := ".html"
	fmt.Println(mime.TypeByExtension(s))

}
