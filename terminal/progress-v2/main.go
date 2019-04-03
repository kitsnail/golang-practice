package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	pb, err := MultiBarNew()
	if err != nil {
		log.Fatal(err)
	}
	total := 100

	pre := fmt.Sprintf("worker %d", 1)
	pb1 := pb.MakeBar(total, pre)
	go pb.Listen()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i <= total; i++ {
			pb1(i)
			time.Sleep(time.Millisecond * 100)
		}
		wg.Done()
	}()
	wg.Wait()
	pb.Println("worker done.")
}
