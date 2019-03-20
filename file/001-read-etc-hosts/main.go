package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	hname := []string{"b[1,5]c[2,3]", "m[2,4,5]"}
	//hname := []string{"m[2,4,5]"}
	hfile := "hosts"
	ips := queryIPbyHost(hfile, hname...)
	fmt.Println(ips)
}

func queryIPbyHost(hfile string, names ...string) (ips []string) {
	filter := regexp.MustCompile("^#")
	var rips []*regexp.Regexp
	for _, name := range names {
		rip := regexp.MustCompile(name)
		rips = append(rips, rip)
	}

	f, err := os.Open(hfile)
	defer f.Close()

	if err != nil {
		return
	}

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if filter.MatchString(line) {
			continue
		}

		for _, rip := range rips {
			if rip.MatchString(line) {
				s := strings.Fields(line)
				ips = append(ips, s[0])
			}
		}
	}
	return
}
