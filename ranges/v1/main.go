package main

import "fmt"

const bs = 2

type rang struct {
	begin int
	end   int
}

func main() {
	totalsize := 11
	var begin int
	var end int

	var ranges []rang
	for begin < totalsize {
		end += bs
		r := rang{begin, end - 1}
		ranges = append(ranges, r)
		begin = end
	}
	for i, r := range ranges {
		fmt.Println(i, r)
	}
}
