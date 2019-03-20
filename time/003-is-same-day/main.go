package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "2017-01-02 15:04:05"
	s2 := "2017-01-02 15:04:05"
	fmt.Println(isSameDate(s1, s2))
}

func isSameDate(s1, s2 string) bool {
	return strings.Fields(s1)[0] == strings.Fields(s2)[0]
}
