package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cname := "rsync"
	args := []string{"-aP", "-e", "ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i testing/paratera_60.id -vvv -p 443 ", "testing/testdata1", "paratera_60@119.90.38.50:/HOME/paratera_60/"}
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
