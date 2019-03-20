package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	host := "192.168.0.130"
	command := "date"
	ret, err := SshRunCmd(host, command)

	if err != nil {
		fmt.Printf("%s", ret)
		os.Exit(1)
	}
	fmt.Printf("ssh run command return: %s", ret)
}

func SshRunCmd(host, command string) ([]byte, error) {
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
	return cmd.CombinedOutput()
}
