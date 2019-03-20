package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
)

func main() {
	writeCsv1()
}

func writeCsv1() {
	var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func writeCsv2() {
	output := "output.csv"
	buf := new(bytes.Buffer)

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(buf)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	fout, err := os.Create(output)
	defer fout.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fout.WriteString(buf.String())
}
