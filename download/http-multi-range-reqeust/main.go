package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const block = int64(204800)

type rang struct {
	begin int64
	end   int64
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("please give an url adress!")
	}
	url := os.Args[1]

	if len(os.Args) < 3 {
		log.Fatalln("please give a save file!")
	}

	file := os.Args[2]

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	download(url, f, block)
}

func download(url string, writer *os.File, block int64) {
	c := &http.Client{}

	hresp, err := c.Head(url)
	if err != nil {
		log.Fatalln(err)
	}
	header := hresp.Header
	conLen := header.Get("Content-Length")
	acceptRange := header.Get("Accept-Ranges")
	clen, err := strconv.ParseInt(conLen, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	if clen < 0 {
		log.Fatalln("requst http resource length:", clen)
	}

	if acceptRange != "bytes" {
		log.Fatalln("request http resource not support part range accept")
	}
	var rangs []rang
	var begin int64
	var end int64
	if block > clen {
		rangs = append(rangs, rang{begin, clen})
	}
	for begin < clen {
		end = begin + block
		rangs = append(rangs, rang{begin, end - 1})
		begin += block
	}
	rangs = append(rangs, rang{end, clen})
	fchan := make(chan struct{}, 30)
	for i, rang := range rangs {
		fmt.Println("start process part:", i)
		rb := fmt.Sprintf("bytes=%d-%d", rang.begin, rang.end)
		go downloadRange(fchan, c, url, rb, writer, rang.begin)
	}
	for _, _ = range rangs {
		<-fchan
	}
	close(fchan)
}

func downloadRange(fchan chan struct{}, c *http.Client, url, rang string, writer *os.File, offset int64) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Range", rang)
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	n, err := writer.WriteAt(b, offset)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("write===>bytes", n)
	fchan <- struct{}{}
}
