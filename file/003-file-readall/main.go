package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "/etc/hosts"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
