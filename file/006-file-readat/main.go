package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const BS = 512

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

	for {
		start = end + 1
		end = start + BS - 1

		if end > size-1 {
			end = size - 1
		}

		b := make([]byte, BS)
		n, err := f.ReadAt(b, start)
		fmt.Println(n)
		fmt.Println(string(b[:n]))
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

}
