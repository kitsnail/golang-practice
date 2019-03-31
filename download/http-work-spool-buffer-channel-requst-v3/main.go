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

const (
	rangeSizeDefault  = int64(10485760)
	bufferSizeDefault = int64(10240)
	threadSizeDefault = 40
)

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
	dlr := NewDownloader(url, f)
	dlr.Download()
}

type Rang struct {
	Begin int64
	End   int64
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

type Downloader struct {
	ThreadSize int
	RangeSize  int64
	BufferSize int64
	ranges     []Rang
	httpClient *http.Client
	completed  *sync.WaitGroup
	allocChan  chan struct{}

	url    string
	writer *os.File
}

func NewDownloader(url string, f *os.File) *Downloader {
	return &Downloader{
		ThreadSize: threadSizeDefault,
		RangeSize:  rangeSizeDefault,
		BufferSize: bufferSizeDefault,
		ranges:     []Rang{},
		httpClient: &http.Client{},
		completed:  &sync.WaitGroup{},
		allocChan:  make(chan struct{}, threadSizeDefault),

		url:    url,
		writer: f,
	}
}

func (d *Downloader) SetRanges(ranges []Rang) {
	d.ranges = ranges
}

func (d *Downloader) SetBufferSize(bs int64) {
	d.BufferSize = bs
}

func (d *Downloader) SetThreadSize(ts int) {
	d.ThreadSize = ts
}

func (d *Downloader) Download() {
	resp, err := d.httpClient.Head(d.url)
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

	ranges := createRanges(totalSize, d.RangeSize)
	d.SetRanges(ranges)
	d.allocate()
}

func (d *Downloader) allocate() {
	for i, _ := range d.ranges {
		d.completed.Add(1)
		go d.downloadRange(i)
	}
	d.completed.Wait()
}

func (d *Downloader) downloadRange(id int) {
	defer d.completed.Done()
	d.allocChan <- struct{}{}
	defer func() { <-d.allocChan }()

	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		log.Println(err)
	}

	rang := d.ranges[id]
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rang.Begin, rang.End))
	resp, err := d.httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	offset := rang.Begin
	p := make([]byte, d.BufferSize)
	for {
		bs, err := resp.Body.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		_, err = d.writer.WriteAt(p, offset)
		if err != nil {
			log.Println(err)
		}
		offset += int64(bs)
	}
}
