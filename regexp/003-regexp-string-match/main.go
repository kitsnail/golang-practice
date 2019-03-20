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
	node := "m1"
	hf := "hosts"
	ips := queryIPbyHostname(hf, node)
	fmt.Println(ips)
}

func queryIPbyHostname(hfile string, name string) (ip string) {
	filter := regexp.MustCompile("^#")
	rip := regexp.MustCompile(name)

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
			//fmt.Println("filter:", line)
			continue

		}

		if rip.MatchString(line) {
			s := strings.Fields(line)
			if strings.Compare(s[1], name) == 0 {
				ip = s[0]
			}
		}
	}

	return
}

func isElemInSlice(slice []string, elem string) bool {
	smap := make(map[string]bool)
	for _, s := range slice {
		smap[s] = true
	}
	return smap[elem]
}
