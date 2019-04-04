package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	debug   bool
	verbose bool
	version bool
	create  string
)

func init() {
	flag.StringVar(&create, "c", "test", "create some things")
	flag.BoolVar(&debug, "D", false, "show debug mode")
	flag.BoolVar(&verbose, "v", false, "show verbose mode")
	flag.BoolVar(&version, "V", false, "show tool version")
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Go is a tool for managing Go source code.\n\n")
		fmt.Printf("Usage:\n\n")
		fmt.Printf("\t go <command> [arguments]\n\n")
		fmt.Printf("Options:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
}
