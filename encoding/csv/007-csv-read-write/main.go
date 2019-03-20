package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "para-input.csv"

	f, err := os.Open(input)
	checkError("Cannot open file", err)
	defer f.Close()
	r := csv.NewReader(f)
	keys, err := r.Read()
	checkError("Cannot read file", err)
	for i, key := range keys {
		fmt.Println(i, ":", key)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
