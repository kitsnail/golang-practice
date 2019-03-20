package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	user := "root"
	host := "192.168.0.28"
	command := "date"
	ret, err := SshRunCmd(user, host, command)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("ssh run command return: %s", ret)
}

func SshRunCmd(user, host, command string) (string, error) {
	addr := fmt.Sprintf("%s@%s", user, host)
	cmd := exec.Command("ssh", addr, command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
