package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	pb := NewProgressBar()
	urls := []string{"http://66.42.91.35/a.txt", "http://66.42.91.35/b.txt", "http://66.42.91.35/c.txt", "http://66.42.91.35/d.txt", "http://66.42.91.35/e.txt"}
	//urls := []string{"http://66.42.91.35/a.txt"}

	var wg sync.WaitGroup
	mch := make(map[int]chan bool)
	for i, url := range urls {
		wg.Add(1)
		nch := make(chan int64)
		compch := make(chan bool)
		mch[i] = compch
		go func(compch chan bool, nch chan int64, id int, url string) {
			Download(id, url, compch, nch)
		}(compch, nch, i, url)

		go func(i int, nch chan int64, ch chan bool) {
			pb.start(i, nch)
			ch <- true
		}(i, nch, compch)
	}

	for _, ch := range mch {
		pb.stop(&wg, ch)
	}

	wg.Wait()
}

func Download(id int, url string, stop chan bool, nch chan int64) {
	ranges := []string{"1-99", "100-199", "200-299", "300-399", "400-499", "500-599", "600-699", "700-799", "800-899", "900-1000"}
	dl := NewDownloader(url, ranges)

	var wg sync.WaitGroup
	tz := 2
	tch := make(chan struct{}, tz)
	for id, _ := range ranges {
		wg.Add(1)
		go func(wg *sync.WaitGroup, id int) {
			defer func() { <-tch }()
			defer wg.Done()
			tch <- struct{}{}
			dl.downloaderRange(id)
		}(&wg, id)
	}

	fch := make(chan bool)
	go func(ch chan bool) {
		wg.Wait()
		ch <- true
	}(fch)
	for {
		select {
		case <-fch:
			stop <- true
			return
		default:
			nch <- dl.completed
			time.Sleep(time.Second)
		}
	}
}

func (p *ProgressBar) start(id int, nch chan int64) {
	p.bufSpace()
	for {
		select {
		case completing, ok := <-nch:
			if !ok {
				return
			}
			p.Printf(id, "[%d] progress bar test %d\n", id, completing)
		default:
			time.Sleep(time.Second / 10)
		}
	}
}

func (p *ProgressBar) stop(wg *sync.WaitGroup, compch chan bool) {
	for {
		select {
		case <-compch:
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			p.Flush()
		}
	}
}
func progress1(url string, completed int64, stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Printf("\r %s ------- %d\n", url, completed)
			return
		default:
			fmt.Printf("\r %s ------- %d %v", url, completed, time.Now().Format(time.Stamp))
			time.Sleep(10 * time.Millisecond)
		}
	}
}

type Downloader struct {
	url       string
	ranges    []string
	completed int64
}

func NewDownloader(url string, ranges []string) *Downloader {
	return &Downloader{
		url:       url,
		ranges:    ranges,
		completed: 0,
	}
}

func (d *Downloader) downloaderRange(id int) {
	atomic.AddInt64(&d.completed, 10)
	time.Sleep(time.Second)
}
