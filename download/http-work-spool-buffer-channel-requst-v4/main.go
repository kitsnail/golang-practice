package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
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

	dlr, err := NewDownloader(url, file)
	if err != nil {
		log.Fatal(err)
	}
	defer dlr.writer.Close()
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

type status struct {
	speed     int64
	completed int64
}

type Downloader struct {
	ThreadSize   int
	RangeSize    int64
	BufferSize   int64
	totalSize    int64
	ranges       []Rang
	status       *status
	httpClient   *http.Client
	completed    *sync.WaitGroup
	allocChan    chan struct{}
	finished     chan bool
	showProgress bool

	filename string
	url      string
	writer   *os.File
}

func NewDownloader(url string, name string) (*Downloader, error) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &Downloader{
		ThreadSize:   threadSizeDefault,
		RangeSize:    rangeSizeDefault,
		BufferSize:   bufferSizeDefault,
		ranges:       []Rang{},
		status:       &status{},
		httpClient:   &http.Client{},
		completed:    &sync.WaitGroup{},
		allocChan:    make(chan struct{}, threadSizeDefault),
		finished:     make(chan bool),
		showProgress: true,

		filename: name,
		url:      url,
		writer:   f,
	}, nil
}

func (d *Downloader) SetTotalSize(n int64) {
	d.totalSize = n
}

func (d *Downloader) SetRanges(ranges []Rang) {
	d.ranges = ranges
}

func (d *Downloader) SetFileName(name string) {
	d.filename = name
}

func (d *Downloader) SetBufferSize(bs int64) {
	d.BufferSize = bs
}

func (d *Downloader) SetThreadSize(ts int) {
	d.ThreadSize = ts
}

func (s *status) Speed() int64 {
	s.speed = s.completed - s.speed
	return s.speed
}

func (d *Downloader) Progress() {
	var speed int64
	for {
		speed = d.status.Speed()
		select {
		case <-d.finished:
			fmt.Printf("\r⇩ %s 100%% %d/%d %d\n", d.filename, d.totalSize, d.totalSize, speed)
			return
		default:
			progress := float64(d.status.completed) / float64(d.totalSize) * 100
			fmt.Printf("\r⇩ %s %.2f%% %d/%d %d", d.filename, progress, d.status.completed, d.totalSize, speed)
			time.Sleep(1 * time.Second)
		}
	}
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
	d.SetTotalSize(totalSize)
	acceptRange := header.Get("Accept-Ranges")
	if acceptRange != "bytes" {
		log.Fatalln("the Request http not support accept ranges")
	}

	ranges := createRanges(d.totalSize, d.RangeSize)
	d.SetRanges(ranges)
	go d.allocate()
	switch {
	case d.showProgress:
		d.Progress()
	default:
		d.Stop()
	}
}

func (d *Downloader) Stop() {
	<-d.finished
}

func (d *Downloader) allocate() {
	for i, _ := range d.ranges {
		d.completed.Add(1)
		go func(id int) {
			if err := d.DownloadRange(id); err != nil {
				d.DownloadRange(id)
			}
		}(i)
	}
	d.completed.Wait()
	d.finished <- true
}

func (d *Downloader) DownloadRange(id int) error {
	defer d.completed.Done()
	d.allocChan <- struct{}{}
	defer func() { <-d.allocChan }()

	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}

	rang := d.ranges[id]
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rang.Begin, rang.End))
	resp, err := d.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	offset := rang.Begin
	p := make([]byte, d.BufferSize)
	var wlock sync.RWMutex
	for {
		bs, err := resp.Body.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		_, err = d.writer.WriteAt(p, offset)
		if err != nil {
			return err
		}
		offset += int64(bs)
		wlock.Lock()
		d.status.completed += int64(bs)
		wlock.Unlock()
	}
	return nil
}
