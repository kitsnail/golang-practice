package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

func main() {
	ranges := []string{"1-99", "100-199", "200-299", "300-399", "400-499", "500-599", "600-699", "700-799", "800-899", "900-1000"}
	tz := 2
	url := "http://66.42.91.35:8001/"
	dl := NewDownloader(url, ranges)
	var wg sync.WaitGroup
	pool, _ := ants.NewPoolWithFunc(tz, func(i interface{}) {
		n := i.(int)
		dl.downloaderRange(n)
		wg.Done()
	})
	defer pool.Release()

	for id, _ := range ranges {
		wg.Add(1)
		pool.Invoke(id)
	}
	wg.Wait()
}

type Downloader struct {
	url    string
	ranges []string
}

func NewDownloader(url string, ranges []string) *Downloader {
	return &Downloader{
		url:    url,
		ranges: ranges,
	}
}

func (d *Downloader) downloaderRange(id int) {
	fmt.Println(d.url, id, d.ranges[id])
	time.Sleep(time.Second)
}
