package main

import (
	"fmt"
	"os"
	"sync"
)

type chunk struct {
	bufsize int
	offset  int64
}

func main() {
	const BufferSize = 100
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fsize := int(fi.Size())

	// Number of go routunes we need to spawn.
	concurrency := fsize / BufferSize

	// All buffer size are the same in the normal case. Offsets depend on the
	// index, Second go routine should start at 100, for example, given our
	// buffer size of 100.
	chunksizes := make([]chunk, concurrency)
	for i := 0; i < concurrency; i++ {
		chunksizes[i].bufsize = BufferSize
		chunksizes[i].offset = int64(BufferSize * i)
	}
	// check for any left over bytes, Add the residual number of bytes as the
	// last chunk size.
	if remainder := fsize % BufferSize; remainder != 0 {
		c := chunk{bufsize: remainder, offset: int64(concurrency * BufferSize)}
		concurrency++
		chunksizes = append(chunksizes, c)
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunksizes []chunk, i int) {
			defer wg.Done()

			chunk := chunksizes[i]
			buffer := make([]byte, chunk.bufsize)
			bytesread, err := f.ReadAt(buffer, chunk.offset)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("bytes read, string(bytestream):", bytesread)
			fmt.Println("bytestream to string", string(buffer))
		}(chunksizes, i)
	}
	wg.Wait()
}
