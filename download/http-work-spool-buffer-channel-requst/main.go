package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const BlockSize = int64(10485760)

type Job struct {
	Id        int
	Range     Rang
	WriteSize int
}

type Result struct {
	Job
}

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

func allocate(jobs chan Job, c *http.Client, f *os.File, url string, ranges []Rang) {
	for i, rang := range ranges {
		n := downloadRange(c, f, url, rang)
		job := Job{i, rang, n}
		jobs <- job
	}
	close(jobs)
}

func download(url string, f *os.File, bs int64) {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	c := &http.Client{}

	resp, err := c.Head(url)
	if err != nil {
		log.Fatalln(err)
	}
	header := resp.Header
	clength := header.Get("Content-Length")
	clen, err := strconv.ParseInt(clength, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	acceptRange := header.Get("Accept-Ranges")
	if acceptRange != "bytes" {
		log.Fatalln("the Request http not support accept ranges")
	}

	var ranges []Rang
	begin := int64(0)
	end := int64(0)
	if bs > clen {
		rang := Rang{begin, clen - 1}
		ranges = append(ranges, rang)
	}
	for begin < clen {
		end += bs
		rang := Rang{begin, end - 1}
		ranges = append(ranges, rang)
		begin = end
	}

	go allocate(jobs, c, f, url, ranges)
	done := make(chan bool)
	go result(done, results)
	noOfWorkers := 10
	createWorkerPool(jobs, results, noOfWorkers)
	<-done
}

func worker(jobs chan Job, results chan Result, wg *sync.WaitGroup) {
	for job := range jobs {
		results <- Result{job}
	}
	wg.Done()
}

func createWorkerPool(jobs chan Job, results chan Result, noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	wg.Wait()
	close(results)
}

func result(done chan bool, results chan Result) {
	for result := range results {
		fmt.Printf("Job id %d, Request bytes=%d-%d ,write bytes %d\n", result.Id, result.Range.Begin, result.Range.End, result.WriteSize)
	}
	done <- true
}

func downloadRange(c *http.Client, f *os.File, url string, rang Rang) (n int) {
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	n, err = f.WriteAt(b, rang.Begin)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}
