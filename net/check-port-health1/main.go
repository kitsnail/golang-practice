package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	hostName := "172.16.1.1"
	portNum := "22"
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second

	conn, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

	if err != nil {
		log.Fatalf("conn DialTimeout error: %s", err.Error())
	}

	buf := make([]byte, 64)

	t1 := time.Now()
	t2 := t1.Add(timeOut)
	conn.SetReadDeadline(t2)
	conn.Read(buf)

	fmt.Printf("read buffer: %s \n", string(buf))
	fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", hostName, int64(timeOut/(1000*time.Millisecond)))
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

}
