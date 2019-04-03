package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/sethgrid/curse"
)

func main() {
	fmt.Println("Progress Bar")
	total := 150
	progressBarWidth := 80
	c, _ := curse.New()

	// give some buffer space on the terminal
	fmt.Println()

	// display a progress bar
	var wg sync.WaitGroup
	for k := 0; k < 3; k++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i <= total; i++ {
				c.MoveUp(1)
				c.EraseCurrentLine()
				fmt.Printf("%d/%d ", i, total)

				c.MoveDown(1)
				c.EraseCurrentLine()
				fmt.Printf("%d %s", id, progressBar(i, total, progressBarWidth))

				time.Sleep(time.Millisecond * 25)
			}
			// end the previous last line of output
			fmt.Println()
			fmt.Println("Complete")
		}(k)
	}
	wg.Wait()
}

func progressBar(progress, total, width int) string {
	bar := make([]string, width)
	for i, _ := range bar {
		if float32(progress)/float32(total) > float32(i)/float32(width) {
			bar[i] = "*"
		} else {
			bar[i] = " "
		}
	}
	return "[" + strings.Join(bar, "") + "]"
}
