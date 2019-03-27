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
	if len(os.Args) < 3 {
		log.Fatalln("please give a save file path")
	}
	url := os.Args[1]
	file := os.Args[2]
	fmt.Println("url:", url)
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	download(url, f)
}

func download(url string, writer io.Writer) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	n, err := io.Copy(writer, resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}
