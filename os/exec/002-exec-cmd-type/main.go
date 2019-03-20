package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	fmt.Printf("cmd.Path: %s\n", cmd.Path)
	fmt.Printf("cmd.Args: %v\n", cmd.Args)
	fmt.Printf("cmd.Env: %v\n", cmd.Env)
	fmt.Printf("cmd.Dir: %s\n", cmd.Dir)
	fmt.Printf("cmd.ExtraFiles: %v\n", cmd.ExtraFiles)
	fmt.Printf("cmd.SysProcAttr: %v\n", cmd.SysProcAttr)
	fmt.Printf("cmd.Process: %v\n", cmd.Process)
	fmt.Printf("cmd.ProcessState: %v\n", cmd.ProcessState)
}
