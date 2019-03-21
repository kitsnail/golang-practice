package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const BufferSize = 100
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesread, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
}
