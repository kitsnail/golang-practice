package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := "localhost:8000"
	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		conn.ReadFrom(os.Stdout)
		log.Println("done")
		done <- struct{}{}
	}()

	go func() {
		conn.Write([]byte("Hello"))
		conn.CloseWrite()
	}()

	<-done
}

func mustCompy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
