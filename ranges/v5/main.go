package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please give a number.")
	}
	ts, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	bs := int64(1024)
	ranges := CreateRanges(ts, bs)

	com := make(chan struct{}, 3)
	go allocate(com, ranges)
	progress(com)
}

func allocate(ach chan<- struct{}, ranges []Rang) {
	var wg sync.WaitGroup
	for i, r := range ranges {
		wg.Add(1)
		go show(&wg, ach, i, r)
	}
	wg.Wait()
}

func show(wg *sync.WaitGroup, ach chan<- struct{}, n int, rang Rang) {
	defer wg.Done()
	defer func() { ach <- struct{}{} }()

	fmt.Printf("Id: %d range %d-%d\n", n, rang.Begin, rang.End)
	time.Sleep(1000 * time.Millisecond)
}

func progress(rch <-chan struct{}) {
	count := 0
	for {
		if _, ok := <-rch; !ok {
			return
		}
		count++
		fmt.Printf("progress %d\n", count)
	}
}

type Rang struct {
	Begin int64
	End   int64
}

func CreateRanges(totalSize int64, bs int64) (ranges []Rang) {
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
