package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("please give a server address")
	}

	checkServerPort(os.Args[1])
}

func checkServerPort(addr string) {
	timeout := time.Duration(5 * time.Second)
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", addr, timeout)
	fmt.Println("wast time:", time.Now().Sub(t1))
	if err != nil {
		log.Println("site unreachable, error:", err)
		return
	}
	fmt.Println("tcp server is ok")
}
