package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "gZ"
	str2 := "GuangzhoU"

	fmt.Println("str1:", str1)
	fmt.Println("str1 -> lower:", strings.ToLower(str1))
	fmt.Println("str2:", str2)
	fmt.Println("str2 -> lower:", strings.ToLower(str2))
}
