package main

import (
	"fmt"
	"path"
)

func main() {
	url := `https://ip62819793.ahcdn.com/key\=fEhYHSk9Jk9wHMHxKP-2rQ,s\=,end\=1554103918,ip\=66\;223/state\=30g+/buffer\=6793000:6793000,6725.0/speed\=76041/reftag\=057661800/42/121/5/127260545/c1/videos/132000/132984/132984_hq.mp4`

	fmt.Println("url", url)
	fmt.Println("url base name", path.Base(url))

}
