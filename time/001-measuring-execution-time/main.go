package main

import (
	"log"
	"time"
)

func main() {
	LongRunningFunction()
}

func LongRunningFunction() {
	defer TimeTaken(time.Now(), "LongRunningFunction")
	time.Sleep(2 * time.Second)
}

func TimeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s took %s\n", name, elapsed)
}
