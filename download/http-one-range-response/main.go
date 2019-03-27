package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const block = 4096

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("please give an url address")
	}
	url := os.Args[1]
	Download(url, os.Stdout, block)
}

func Download(url string, writer io.Writer, block int) {
	c := http.Client{}

	hresp, err := c.Head(url)
	if err != nil {
		log.Fatalln(err)
	}
	header := hresp.Header
	contentLen := header.Get("Content-Length")
	acceptRange := header.Get("Accept-Ranges")
	clen, err := strconv.Atoi(contentLen)
	if err != nil {
		log.Fatalln(err)
	}
	if clen < 0 {
		log.Fatalln("request source body lenth:", contentLen)
	}
	if acceptRange != "bytes" {
		log.Fatalln("request source not support Accept range bytes!")
	}

	breq, err := http.NewRequest("GET", url, nil)
	var rang string
	if block > clen {
		rang = fmt.Sprintf("bytes=0-%d", clen)
	} else {
		rang = fmt.Sprintf("bytes=0-%d", block)
	}
	breq.Header.Set("Range", rang)
	bresp, err := c.Do(breq)
	if err != nil {
		log.Fatalln(err)
	}
	defer bresp.Body.Close()
	io.Copy(writer, bresp.Body)
}
