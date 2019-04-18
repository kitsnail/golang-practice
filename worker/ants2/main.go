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
	pool, _ := ants.NewPool(tz)
	defer pool.Release()
	var wg sync.WaitGroup
	for id, rang := range ranges {
		wg.Add(1)
		dlr := NewDownloader(id, url, rang)
		dlFunc := func() {
			dlr.downloaderRange()
			wg.Done()
		}
		pool.Submit(dlFunc)
	}
	wg.Wait()
}

type Downloader struct {
	id   int
	url  string
	rang string
}

func NewDownloader(id int, url string, rang string) *Downloader {
	return &Downloader{
		id:   id,
		url:  url,
		rang: rang,
	}
}

func (d *Downloader) downloaderRange() {
	fmt.Println(d.url, d.id, d.rang)
	time.Sleep(time.Second)
}
