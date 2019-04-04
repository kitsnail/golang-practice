package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan Rang, results chan<- int64) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j.Begin, j.End)
		time.Sleep(2 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- 2
	}
}

type Rang struct {
	Begin int64
	End   int64
}

func createRanges(ts, bs int64) (ranges []Rang) {
	var begin int64
	var end int64
	for begin < ts {
		end += bs
		ranges = append(ranges, Rang{begin, end})
		begin = end
	}
	ranges[len(ranges)-1].End = ts - 1
	return
}

func main() {
	jobs := make(chan Rang, 2)
	results := make(chan int64, 2)

	ranges := createRanges(50, 3)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for _, r := range ranges {
		jobs <- r
	}
	close(jobs)

	for i, r := range ranges {
		fmt.Println("======", i, r)
		<-results
	}

}
