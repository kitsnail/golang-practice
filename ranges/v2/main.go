package main

import "fmt"

type Rang struct {
	Begin int
	End   int
}

func main() {
	ts := 330938
	bs := 1024

	var begin int
	var end int
	var ranges []Rang
	for begin < ts {
		end += bs
		ranges = append(ranges, Rang{begin, end - 1})
		begin = end
	}
	ranges[len(ranges)-1].End = ts - 1

	for i, r := range ranges {
		fmt.Println(i, r)
	}
}
