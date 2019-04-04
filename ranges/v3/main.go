package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	for i, r := range ranges {
		fmt.Println(i, r)
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
