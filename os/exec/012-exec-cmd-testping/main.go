package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	host := "192.168.0.123"
	if testPing(host) {
		fmt.Printf("ping %s is True.\n", host)
	} else {
		fmt.Printf("ping %s is False.\n", host)
	}
}

// testPing ping test host network is well
func testPing(host string) bool {
	// set ping run command args
	pingArgs := fmt.Sprintf("ping -c1 -w1 %s | grep 'packet' | awk '{print$6}'", host)
	cmd := exec.Command("sh", "-c", pingArgs)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return false
	}
	if strings.EqualFold("0%", strings.TrimSpace(stdout.String())) {
		return true
	}
	return false
}
