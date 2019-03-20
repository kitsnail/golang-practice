package main

import (
	"fmt"
)

func main() {
	tmap := make(map[string]string)
	tmap["tid001"] = "stop"
	value, ok := tmap["tid001"]
	fmt.Println("ok:", ok)
	fmt.Println("value:", value)
}
