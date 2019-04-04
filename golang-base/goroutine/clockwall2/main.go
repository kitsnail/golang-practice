// clockwall listens to multiple clock servers concurrently.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

type clock struct {
	name, host string
}

//func (c *clock) watch(w io.Writer, r io.Reader) {
func (c *clock) watch(up int, r io.Reader) {

	s := bufio.NewScanner(r)
	tm.Clear() // Clear current screen
	for s.Scan() {
		tm.Printf("%s\t%s\n", c.name, s.Text())
		tm.MoveCursorUp(up)
		tm.Flush()
		//fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST ...")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, a := range os.Args[1:] {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", a)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}
	for n, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		//go c.watch(os.Stdout, conn)
		go c.watch(n, conn)
	}
	// Sleep while other goroutines do the work.
	for {
		time.Sleep(time.Minute)
	}
}
