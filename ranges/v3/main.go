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
	bs := int64(10)
	ranges := CreateRanges(5, bs, ts)

	for i, r := range ranges {
		fmt.Println(i, r)
	}
}

type Rang struct {
	Begin int64
	End   int64
}

func CreateRanges(begin, bs, total int64) (ranges []Rang) {
	var end int64

	for begin < total {
		end = bs + begin
		ranges = append(ranges, Rang{begin, end})
		begin = end
	}
	ranges[len(ranges)-1].End = total - 1
	return
}
