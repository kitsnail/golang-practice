package main

import (
	"fmt"
	"os"

	"encoding/json"
)

func main() {
	// Perm defualt 0644
	file1 := "test.data"

	// show the file stat
	fi, err := os.Stat(file1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("file name:", fi.Name())
	fmt.Println("file size:", fi.Size())
	fmt.Println("file mode:", fi.Mode())
	fmt.Println("file mod time:", fi.ModTime())
	fmt.Println("file sys:", fi.Sys())

	// file Sys marshal to json
	fs, err := json.Marshal(fi.Sys())
	if err != nil {
		fmt.Println("json Marshal error:", err)
	}
	fmt.Println("file sys json:", string(fs))
}
