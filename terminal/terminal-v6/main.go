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
	for i := 0; i < 30; i++ {
		pb.bufSpace()
		go func(id int) {
			for k := 0; k < 10000; k++ {
				pb.Printf(id, "[%d] progress bar test %d\n", id, k)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		pb.Flush()
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
