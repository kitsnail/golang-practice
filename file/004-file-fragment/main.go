package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const BodySize = 1 * 1024

func main() {
	file := "testfile"

	fi, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	size := fi.Size()

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conc := make(chan struct{}, 5)

	var (
		start int64 = -1
		end   int64 = -1
	)

	var wg sync.WaitGroup

	for {
		start = end + 1
		end = start + BodySize - 1

		if end > size-1 {
			end = size - 1
		}

		var buf bytes.Buffer
		_, err := f.Seek(start, 0)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.CopyN(&buf, f, end-start+1)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		conc <- struct{}{}
		go upload(&wg, conc, fmt.Sprintf("%d-%d/%d\n %s", start, end, size, buf.String()))

		if end >= size-1 {
			break
		}
	}
	wg.Wait()
}

func upload(wg *sync.WaitGroup, conc chan struct{}, body string) {
	defer wg.Done()
	defer func() {
		<-conc
	}()

	fmt.Printf("%s\n", body)
	fmt.Println("=================================")
}
