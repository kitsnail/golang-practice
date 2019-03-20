package main

import (
	"fmt"
	"os"
)

func main() {
	pubkey := "pub/pubkey.id"
	err := CheckPubkeyPerm(pubkey)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("check perm is good.")
}

func CheckPubkeyPerm(name string) error {
	var pubkeyPerm os.FileMode = 0600
	fi, err := os.Stat(name)
	if err != nil {
		return err
	}

	if fi.Mode() != pubkeyPerm {
		if err := os.Chmod(name, pubkeyPerm); err != nil {
			return err
		}
	}
	return nil
}
