package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const BS = 256

func main() {
	file := "testfile"

	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal("1", err)
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size := fi.Size()

	var (
		start int64 = -1
		end   int64 = -1
	)
	conc := make(chan struct{}, 5)
	var wg sync.WaitGroup

	for {
		start = end + 1
		end = start + BS - 1

		if end > size-1 {
			end = size - 1
		}

		b := make([]byte, BS)
		n, err := f.ReadAt(b, start)
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		conc <- struct{}{}
		go upload(&wg, conc, fmt.Sprintf("%d-%d/%d\n %s", start, end, size, string(b[:n])))
		if n == 0 || err == io.EOF {
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

	fmt.Println(body)
	fmt.Println("==========")
}
