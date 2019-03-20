package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cname := "ssh1"
	args := []string{"root@vm.centos7", "-p", "2222"}
	err := cmdRun(cname, args)
	if err != nil {
		fmt.Println("main.err:", err.Error())
	}
}

func cmdRun(name string, args []string) (err error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
