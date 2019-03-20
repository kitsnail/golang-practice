package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	csvName := "Export_parafile_201707251216_xN.csv"
	file, err := os.Open(csvName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(record) // record has the type []string
	}
}
