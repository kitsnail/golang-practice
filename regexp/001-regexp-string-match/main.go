package main

import (
	"fmt"
	"regexp"
)

func main() {
	nodes := "dm"
	line := "192.168.0.111 dm"
	reg := regexp.MustCompile(nodes)
	if reg.MatchString(line) {
		fmt.Println("ture")
	} else {
		fmt.Println("false")
	}
}
