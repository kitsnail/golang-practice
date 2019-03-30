package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const BlockSize = int64(10485760)
const ThreadSize = 40

type Rang struct {
	Begin int64
	End   int64
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please give an URL string!")
	}
	url := os.Args[1]

	if len(os.Args) < 3 {
		log.Fatalln("Please give a save file path")
	}
	file := os.Args[2]

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	download(url, f, BlockSize)
}

func createRanges(totalSize int64, bs int64) (ranges []Rang) {
	var begin int64
	var end int64

	for begin < totalSize {
		end += bs
		ranges = append(ranges, Rang{begin, end})
		begin = end
	}
	ranges[len(ranges)-1].End = totalSize - 1
	return
}

func download(url string, f *os.File, bs int64) {
	c := &http.Client{}

	resp, err := c.Head(url)
	if err != nil {
		log.Fatalln(err)
	}
	header := resp.Header
	clength := header.Get("Content-Length")
	totalSize, err := strconv.ParseInt(clength, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	acceptRange := header.Get("Accept-Ranges")
	if acceptRange != "bytes" {
		log.Fatalln("the Request http not support accept ranges")
	}

	ranges := createRanges(totalSize, BlockSize)
	jobs := make(chan struct{}, ThreadSize)
	allocate(jobs, c, f, url, ranges)
}

func allocate(jobs chan struct{}, c *http.Client, f *os.File, url string, ranges []Rang) {
	var wg sync.WaitGroup
	for _, rang := range ranges {
		wg.Add(1)
		go downloadRange(&wg, jobs, c, f, url, rang)
	}
	wg.Wait()
}

/*
func result(done chan bool, results chan Result) {
	for result := range results {
		fmt.Printf("Job id %d, Request bytes=%d-%d ,write bytes %d\n", result.Id, result.Range.Begin, result.Range.End, result.WriteSize)
	}
	done <- true
}
*/

func downloadRange(wg *sync.WaitGroup, jobs chan struct{}, c *http.Client, f *os.File, url string, rang Rang) {
	defer wg.Done()
	jobs <- struct{}{}
	defer func() { <-jobs }()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rang.Begin, rang.End))
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	offset := rang.Begin
	p := make([]byte, 1024)
	for {
		bs, err := resp.Body.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		_, err = f.WriteAt(p, offset)
		if err != nil {
			log.Fatalln(err)
		}
		offset += int64(bs)
	}
}
