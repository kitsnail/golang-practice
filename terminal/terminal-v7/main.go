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
	compch := make(chan bool)
	for i, total := range totals {
		go func(i int, total int, ch chan bool) {
			pb.start(i, total)
			ch <- true
		}(i, total, compch)
	}
	for _, _ = range totals {
		pb.stop(compch)
	}
}

func (p *ProgressBar) start(id int, total int) {
	p.bufSpace()
	for k := 0; k < total; k++ {
		p.Printf(id, "[%d] progress bar test %d\n", id, k)
		time.Sleep(time.Second / 10)
	}
}

func (p *ProgressBar) stop(compch chan bool) {
	for {
		select {
		case <-compch:
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
