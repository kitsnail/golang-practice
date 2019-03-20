package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	host := "192.168.0.128"
	command := "date"
	ret, err := SshRunCmd(host, command)
	if err != nil {
		fmt.Printf("%s", ret)
		os.Exit(1)
	}
	fmt.Printf("ssh run command return: %s", ret)
}

func SshRunCmd(host, command string) (string, error) {
	// set ssh run command args
	sshArgs := []string{
		"-o",
		"PasswordAuthentication=no",
		"-o",
		"StrictHostKeyChecking=no",
		"-o",
		"UserKnownHostsFile=/dev/null",
		"-o", "ConnectTimeout=5"}
	sshArgs = append(sshArgs, host, command)
	cmd := exec.Command("ssh", sshArgs...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}
