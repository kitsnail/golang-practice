package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

var sum int32

func myFun(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
	time.Sleep(time.Second)
}

func demoFunc() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 15

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	p1, _ := ants.NewPool(3)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p1.Submit(syncCalculateSum)
	}
	wg.Wait()

	fmt.Printf("running gotoutine: %d\n", p1.Running())
	fmt.Printf("finish all tasks.\n")

	p, _ := ants.NewPoolWithFunc(4, func(i interface{}) {
		myFun(i)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
