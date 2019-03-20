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
	//nodes := []string{"m[0-9]", "dm", "sdwqd"}
	nodes := []string{"m1"}
	hf := "hosts"
	ips := queryIPbyHost(hf, nodes...)
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
			//fmt.Println("filter:", line)
			continue

		}

		for _, rip := range rips {
			if rip.MatchString(line) {
				s := strings.Fields(line)
				ips = append(ips, s[1])
			}
		}
	}

	errCount := 0

	if len(ips) == 0 {
		for _, name := range names {
			fmt.Printf("ips =0 : \"%s\" is no match in %s\n", name, hfile)
			errCount++
		}
	}

	for _, rip := range rips {
		okcount := 0
		for _, ip := range ips {
			if rip.MatchString(ip) {
				okcount++
			}
		}
		if okcount == 0 {
			fmt.Printf("ips >0 : \"%s\" is no match in %s\n", rip.String(), hfile)
			errCount++
		}
	}

	if errCount > 0 {
		os.Exit(1)
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
