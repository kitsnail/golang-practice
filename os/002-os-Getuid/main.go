package main

import (
	"fmt"
	"os"
)

func main() {
	// show current login user id
	fmt.Println("This current login uid:", os.Getuid())
}
