package main

import "fmt"

func main() {
	mmp := make(map[string][]string)
	e, ok := mmp["aaa"]
	fmt.Println("e:", e)
	fmt.Println("ok:", ok)
}
