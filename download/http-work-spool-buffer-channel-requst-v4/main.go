package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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
	var saveFile string
	flag.StringVar(&saveFile, "s", "", "give a save to file")
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("Please give an URL string!")
	}
	url := flag.Arg(0)

	if saveFile == "" {
		saveFile = path.Base(url)
	}

	dlr, err := NewDownloader(url, saveFile)
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
	completing int64
	completed  int64
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
	speed := s.completing - s.completed
	s.completed = s.completing
	return speed
}

func (d *Downloader) Progress() {
	var speed string
	for {
		speed = BytesToHuman(float64(d.status.Speed())) + "/s"
		select {
		case <-d.finished:
			fmt.Printf("\r⇩ %s 100%% %d/%d %-15s %-13s\n", d.filename, d.totalSize, d.totalSize, speed, "[Completed]")
			return
		default:
			progress := float64(d.status.completed) / float64(d.totalSize) * 100
			fmt.Printf("\r⇩ %s %.2f%% %d/%d %-15s %-13s", d.filename, progress, d.status.completed, d.totalSize, speed, "[InProgess]")
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
				d.completed.Add(1)
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
		d.status.completing += int64(bs)
		wlock.Unlock()
	}
	return nil
}

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// bytesToHuman humen readable bumber
// num unit default Byte
// ret unit B,KB,MB,GB,TB,PB,EB,ZB,YB
// eg. num = 1024
// return 1,KB
func bytesToHuman(num float64) (ret float64, unit string) {
	switch {
	case num >= 0 && num < KB:
		ret = num
		unit = "B"
	case num >= KB && num < MB:
		ret = num / KB
		unit = "KiB"
	case num >= MB && num < GB:
		ret = num / MB
		unit = "MiB"
	case num >= GB && num < TB:
		ret = num / GB
		unit = "GiB"
	case num >= TB && num < PB:
		ret = num / TB
		unit = "TiB"
	case num >= PB && num < EB:
		ret = num / PB
		unit = "PiB"
	case num >= EB && num < ZB:
		ret = num / EB
		unit = "EiB"
	case num >= ZB && num < YB:
		ret = num / ZB
		unit = "ZiB"
	case num >= YB:
		ret = num / YB
		unit = "YiB"
	default:
		ret = num
		unit = "UNKOWN"
	}
	return
}

func BytesToHuman(num float64) (ret string) {
	rn, unit := bytesToHuman(num)
	return fmt.Sprintf("%.2f%s", rn, unit)
}
