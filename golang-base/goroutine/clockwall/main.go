package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please give a or more region clock server adress\n")
		fmt.Println("eg. NewYork=localhost:8000")
		os.Exit(1)
	}
	servers := os.Args[1:]
	stop := make(chan bool)
	for _, server := range servers {
		sli := strings.Split(server, "=")
		if len(sli) < 2 {
			continue
		}
		region := sli[0]
		addr := sli[1]
		go showClock(stop, region, addr)
	}
	<-stop
	close(stop)
}

func showClock(stop chan bool, region string, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
		stop <- true
	}
	defer conn.Close()
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
		stop <- true
	}
	//fmt.Printf("\r %s: %s", region, string(b))
}
