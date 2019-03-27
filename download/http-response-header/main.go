package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("please give a url string")
	}
	url := os.Args[1]
	f := os.Stdout
	download(url, f)
}

func download(url string, writer io.Writer) {
	c := http.Client{}
	resp, err := c.Head(url)
	if err != nil {
		log.Fatalln(err)
	}
	header := resp.Header
	contentLen := header.Get("Content-Length")
	acceptRange := header.Get("Accept-Ranges")
	fmt.Println(contentLen)
	fmt.Println(acceptRange)
}
