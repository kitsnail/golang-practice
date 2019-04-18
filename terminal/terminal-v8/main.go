package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/sethgrid/curse"
)

func main() {
	pb := NewProgressBar()
	totals := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	var wg sync.WaitGroup
	mch := make(map[int]chan bool)
	for i, total := range totals {
		wg.Add(1)
		nch := make(chan int)
		compch := make(chan bool)
		mch[i] = compch
		go working(total, nch)

		go func(i int, nch chan int, ch chan bool) {
			pb.start(i, nch)
			ch <- true
		}(i, nch, compch)
	}

	for _, ch := range mch {
		pb.stop(&wg, ch)
	}

	wg.Wait()
}

func working(total int, ch chan int) {
	for i := 0; i < total; i++ {
		time.Sleep(time.Second / 10)
		ch <- i
	}
	close(ch)
}

func (p *ProgressBar) start(id int, nch chan int) {
	p.bufSpace()
	for {
		select {
		case completing, ok := <-nch:
			if !ok {
				return
			}
			p.Printf(id, "[%d] progress bar test %d\n", id, completing)
		default:
			time.Sleep(time.Second / 10)
		}
	}
}

func (p *ProgressBar) stop(wg *sync.WaitGroup, compch chan bool) {
	for {
		select {
		case <-compch:
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			p.Flush()
		}
	}
}

type ProgressBar struct {
	output *bufio.Writer

	history map[int]string
	mtx     *sync.RWMutex
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		output:  bufio.NewWriter(os.Stdout),
		history: make(map[int]string),
		mtx:     new(sync.RWMutex),
	}
}

func (p *ProgressBar) bufSpace() {
	fmt.Println()
}

func (p *ProgressBar) Flush() {
	c, _ := curse.New()
	total := len(p.history)
	c.MoveUp(total)
	for n := 0; n < total; n++ {
		p.output.WriteString(p.history[n])
	}
	p.output.Flush()
}

func (p *ProgressBar) Print(id int, a ...interface{}) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.history[id] = fmt.Sprint(a...)
	return
}

func (p *ProgressBar) Println(id int, a ...interface{}) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.history[id] = fmt.Sprintln(a...)
	return
}

func (p *ProgressBar) Printf(id int, format string, a ...interface{}) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.history[id] = fmt.Sprintf(format, a...)
	return
}
