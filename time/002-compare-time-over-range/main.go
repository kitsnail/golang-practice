package main

import (
	"fmt"
	"log"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func main() {
	s1 := "2017-07-31 14:18:06"
	s2 := "2017-08-01 18:18:06"
	d := 24

	t1, err := time.Parse(timeLayout, s1)
	if err != nil {
		log.Fatalln(err)
	}
	t2, err := time.Parse(timeLayout, s2)
	if err != nil {
		log.Fatalln(err)
	}

	dr := d * 3600

	fmt.Println(CompareTimeOverRange(t1, t2, float64(dr)))
}

func CompareTimeOverRange(t1, t2 time.Time, d float64) bool {
	dt := t2.Sub(t1).Seconds()
	if dt > d {
		return false
	}
	return true
}
