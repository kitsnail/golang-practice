package main

import (
	"fmt"
	"time"
)

func main() {
	jobch := genJob(10)
	retch := make(chan string, 5)
	workerPool(5, jobch, retch)

	time.Sleep(time.Second)
	close(retch)

	for ret := range retch {
		fmt.Println(ret)
	}
}

func worker(id int, jobch <-chan int, retch chan<- string) {
	for job := range jobch {
		ret := fmt.Sprintf("worker %d processed job: %d", id, job)
		retch <- ret
	}
}

func workerPool(n int, jobch <-chan int, retch chan<- string) {
	for i := 0; i < n; i++ {
		go worker(i, jobch, retch)
	}
}

func genJob(n int) <-chan int {
	jobch := make(chan int, 5)
	go func() {
		for i := 0; i < n; i++ {
			jobch <- i
		}
		close(jobch)
	}()
	return jobch
}
