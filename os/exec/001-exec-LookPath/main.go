package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd1 := "echo"
	cmd2 := "echo1"
	run(cmd1)
	run(cmd2)
}
func run(command string) {
	path, err := exec.LookPath(command)
	if err != nil {
		log.Fatalf("%s is not exists, %s", command, err.Error())
	}
	fmt.Printf("%s is available at %s\n", command, path)
}
