package main

import (
	"fmt"
	"os"
)

func main() {
	// show currrent login user's gid
	fmt.Println("The Login user's GID is:", os.Getgid())
}
