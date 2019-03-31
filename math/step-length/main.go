package main

import (
	"fmt"
	"time"
)

func main() {
	var step int
	for i := 5; i < 100; i += 5 {
		speed := i - step
		step = i
		fmt.Println(i, speed, step)
		time.Sleep(1 * time.Second)
	}
}
